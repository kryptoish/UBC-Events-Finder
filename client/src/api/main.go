package main

import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"

	"github.com/spf13/viper"
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

func main() {
	var token string
	if (local) {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
		token = viper.Get("qAPI_KEY").(string)
	} else {
		token = os.Getenv("qAPI_KEY")
	}

    /*username := os.Getenv("qUSER") /* currently using my personal acct because
                                that is what I was able to set the API
                                access through */

    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/auth/callback", handleAuthRedirect)
    
    post := retrieve_post_data(token, retrieve_user_id(token))
	fmt.Printf("cap: %s", post.Data[0].Caption);

    //fmt.Print("before")
    http.ListenAndServe(":8080", nil)
}

func handleRoot (w http.ResponseWriter, r *http.Request) {
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
    post := retrieve_post_data(token, retrieve_user_id(token))
    
    response := map[string]string{"message": post.Data[0].Caption}
    json.NewEncoder(w).Encode(response)
}

func handleAuthRedirect(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Authorization code not found", http.StatusBadRequest)
        return
    }

    //untested
    token := get_token(code)
    fmt.Printf("Not an error: %s", token)
    /*
    userID := retrieve_user_id(token)
    //err = store_token(userID, token) //something to do with database stuff
                                    // maybe PostgreSQL on Render
    if err != nil {
        http.Error(w, "Failed to store token", http.StatusInternalServerError)
        return
    } 
    */

    http.Redirect(w, r, "/?authStatus=success", http.StatusFound)
}
