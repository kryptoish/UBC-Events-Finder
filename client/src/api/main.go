package main

import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"
)

type MediaResponse struct {
	Data []struct {
		ID        string `json:"id"`
		Caption   string `json:"caption"`
		MediaURL  string `json:"media_url"`
		Permalink string `json:"permalink"`
	} `json:"data"`
}


func main() {
    token := os.Getenv("qAPI_KEY")
    username := os.Getenv("qUSER") /* currently using my personal acct because
                                that is what I was able to set the API
                                access through */

    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/auth/callback", handleAuthRedirect)
    
	post := retrieve_post_data(token, username)
	fmt.Printf("cap: %s", post.Data[0].Caption);
    http.HandleFunc("/ping", pingHandler)

    //fmt.Print("before")
    http.ListenAndServe(":8080", nil)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is alive"))
}

func handleRoot (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    
    token := os.Getenv("qAPI_KEY")
    username := os.Getenv("qUSER") 
	
	//get_token()
    post := retrieve_post_data(token, username)
    
    response := map[string]string{"message": post.Data[0].Caption}
    json.NewEncoder(w).Encode(response)
}

func handleAuthRedirect(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Authorization code not found", http.StatusBadRequest)
        return
    }

    fmt.Printf("Authorization Code: %s\n", code)

    fmt.Fprintf(w, "Authorization successful! Code: %s", code)
}
