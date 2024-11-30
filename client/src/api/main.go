package main

import (
    "os"
    "encoding/json"
    "net/http"

	"github.com/spf13/viper" /* for reading local .env file */
    "github.com/rs/cors" /* for handling re-routing the oauth pipeline */
)

/* Set 'true' when testing locally, 
    when true, data is sent to localhost:8080 */
const local = false;

/* Data recieved from the media call to the instagram graph api */
type MediaResponse struct {
	Data []struct {
		ID        string `json:"id"`
		Caption   string `json:"caption"`
		MediaURL  string `json:"media_url"`
		Permalink string `json:"permalink"`
	} `json:"data"`
}

/* Data to send to the frontend */
type ProcessedResponse struct {
	Data []struct {
		ID        string `json:"id"`
		Caption   string `json:"caption"`
		MediaURL  string `json:"media_url"`
		Permalink string `json:"permalink"`
        Username  string `json:"username"`
        Food      string `json:"food"`
        Date      string `json:"date"`
        Time      string `json:"time"`
        Location  string `json:"location"`
	} `json:"data"`
}

/* Setus up cors mux to allow rerouting, and listens
    on port 8080*/
func main() {
    var frontend string

	if (local) {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
        frontend = "http://localhost:5173"
	} else {
        frontend = "https://ubc-events-finder.vercel.app"
	}
    
    c := cors.New(cors.Options{
        AllowedOrigins: []string{frontend},
        AllowCredentials: true,
    })

    mux := http.NewServeMux()
    mux.HandleFunc("/auth/callback", handleAuthRedirect)
    mux.HandleFunc("/", handleRoot)

    handler := c.Handler(mux)
    address := ":8080"

    http.ListenAndServe(address, handler)
}

/* Handles getting the post data from instagram api,
    processing it, then sending back, every time
    the page is refreshed.                          */
func handleRoot (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    
    /* Currently storing each token in env variable 
        Integrating a database is a future step */
	var chbeToken string
    var eusToken string
    var eceToken string
	if (local) {
		chbeToken = viper.Get("CHBE_KEY").(string)
		eusToken = viper.Get("EUS_KEY").(string)
		eceToken = viper.Get("ECE_KEY").(string)
	} else {
		chbeToken = os.Getenv("CHBE_KEY")
		eusToken = os.Getenv("EUS_KEY")
		eceToken = os.Getenv("ECE_KEY")
	}
   
    var relaventData []ProcessedResponse

    relaventData = append(relaventData, relaventInfo(chbeToken, "chbecouncil", w))
    relaventData = append(relaventData, relaventInfo(eusToken, "ubcengineers", w))
    relaventData = append(relaventData, relaventInfo(eceToken, "eceubc", w))

    finalData := mergedResponses(relaventData)	
    json.NewEncoder(w).Encode(finalData) 
}

/* Handles the getting a long term access token from the authorization code
   when a new account gives us permission to use their account             */
func handleAuthRedirect(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Authorization code not found", http.StatusBadRequest)
        return
    }

    token, err := getToken(code)
    if err != nil {
        http.Error(w, "Authorization code not found", http.StatusBadRequest)
        return
    }
    response := map[string]string{"token": token}
    json.NewEncoder(w).Encode(response)

}
