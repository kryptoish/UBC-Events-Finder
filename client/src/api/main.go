package main

import (
    "fmt"
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


func handleRoot (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    
    /* Currently storing each token in env variable 
        Integrating a database is a future step */
	var chbe_token string
    var eus_token string
    var ece_token string
	if (local) {
		chbe_token = viper.Get("CHBE_KEY").(string)
		eus_token = viper.Get("EUS_KEY").(string)
		ece_token = viper.Get("ECE_KEY").(string)
	} else {
		chbe_token = os.Getenv("CHBE_KEY")
		eus_token = os.Getenv("EUS_KEY")
		ece_token = os.Getenv("ECE_KEY")
	}
   
    var relavent_data []ProcessedResponse

    relavent_data = append(relavent_data, relavent_info(chbe_token, "chbecouncil", w))
    relavent_data = append(relavent_data, relavent_info(eus_token, "ubcengineers", w))
    relavent_data = append(relavent_data, relavent_info(ece_token, "eceubc", w))

    final_data := mergedResponses(relavent_data)	
    json.NewEncoder(w).Encode(final_data) 
}

func handleAuthRedirect(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    if code == "" {
        fmt.Print("no code\n")
        http.Error(w, "Authorization code not found", http.StatusBadRequest)
        return
    }

    //untested
    token, err := get_token(code)
    if err != nil {
        http.Error(w, "Authorization code not found", http.StatusBadRequest)
    }
    response := map[string]string{"token": token}
    json.NewEncoder(w).Encode(response)

}
