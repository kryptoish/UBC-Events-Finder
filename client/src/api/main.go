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
	if (local) {
		chbe_token = viper.Get("CHBE_KEY").(string)
	} else {
		chbe_token = os.Getenv("CHBE_KEY")
	}
   
    fmt.Printf("CHBE_KEY: %s\n", chbe_token)
    var relavent_data []ProcessedResponse

    user_id, username := retrieve_user_id(chbe_token, w)
    fmt.Printf("chbe username: %s\n", username)
    posts := retrieve_post_data(chbe_token, user_id, w)
    fmt.Printf("chbe post: %s\n", posts.Data[0].Caption)
    food_posts := filter_data(posts)
    relavent_data = append(relavent_data, relavent_info(food_posts, username))

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
    /*
    userID := retrieve_user_id(token)
    //err = store_token(userID, token) //something to do with database stuff
                                    // maybe PostgreSQL on Render
    if err != nil {
        http.Error(w, "Failed to store token", http.StatusInternalServerError)
        return
    } 
    */

}
