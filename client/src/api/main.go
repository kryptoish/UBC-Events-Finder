package main

import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"

	"github.com/spf13/viper"
    "github.com/rs/cors"
)

const local = false;

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
	var token string
    var frontend string
	if (local) {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
		token = viper.Get("qAPI_KEY").(string)
        frontend = "http://localhost:5173"
	} else {
		token = os.Getenv("qAPI_KEY")
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

    /*username := os.Getenv("qUSER") /* currently using my personal acct because
                                that is what I was able to set the API
                                access through */
    
    //fmt.Print("before")
    address := ":8080"
    if local {
        address = "localhost:8080"
    }
    fmt.Printf("server running at http://%s\n, token=%s\n\n", address,token)
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
    
	var token string
	if (local) {
		token = viper.Get("qAPI_KEY").(string)
	} else {
		token = os.Getenv("qAPI_KEY")
	}
    //username := os.Getenv("qUSER") 
	
	//get_token()
    posts := retrieve_post_data(token, retrieve_user_id(token))
    food_posts := filter_data(posts)
    
    json.NewEncoder(w).Encode(food_posts)
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
        http.Redirect(w, r, "https://google.com", http.StatusOK)
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
