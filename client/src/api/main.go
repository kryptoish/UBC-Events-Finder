package main

import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"

	"github.com/spf13/viper"
    "github.com/rs/cors"
)

const local = true;

type MediaResponse struct {
	Data []struct {
		ID        string `json:"id"`
		Caption   string `json:"caption"`
		MediaURL  string `json:"media_url"`
		Permalink string `json:"permalink"`
	} `json:"data"`
}

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
    mux.HandleFunc("/ping", pingHandler)
    mux.HandleFunc("/", handleRoot)

    handler := c.Handler(mux)
    address := ":8080"
    if local {
        address = "localhost:8080"
    }
    http.ListenAndServe(address, handler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Print("handle ping\n")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is alive"))
}

func handleRoot (w http.ResponseWriter, r *http.Request) {
    fmt.Print("handle root\n")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    
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

    chbe_id, _ := retrieve_user_id(chbe_token, w)
    chbe_posts := retrieve_post_data(chbe_token, chbe_id, w)
    chbe_food_posts := filter_data(chbe_posts)
    relavent_data = append(relavent_data, relavent_info(chbe_food_posts, "chbecouncil"))
    
    eus_id, _ := retrieve_user_id(eus_token, w)
    eus_posts := retrieve_post_data(eus_token, eus_id, w)
    eus_food_posts := filter_data(eus_posts)
    relavent_data = append(relavent_data, relavent_info(eus_food_posts, "ubcengineers"))
    
    ece_id, _ := retrieve_user_id(ece_token, w)
    ece_posts := retrieve_post_data(ece_token, ece_id, w)
    ece_food_posts := filter_data(ece_posts)
    relavent_data = append(relavent_data, relavent_info(ece_food_posts, "eceubc"))


    final_data := mergedResponses(relavent_data)	
    json.NewEncoder(w).Encode(final_data) 
}

func handleAuthRedirect(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("handle callback at URL: %s\n", r.URL.String())
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