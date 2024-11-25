package main

import (
    "fmt"
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

const token = "IGQWRQMWMyYkxxckFFYXFLZAkhoZATdJeHhKa205MWxtVzAtYzBMaVpYN0JrY1FxU0h3amQ1SkZAMeWs3dTdEQ2R3RUJtenZAfR0tQLVQzVnItWXFBZAzBUUDZAhMWJJd094eTBycDFJSVhQRFIzNkR2N1Bvc2VZAYTJwWmcZD"
const username = "quansenycz" /* currently using my personal acct because
                            that is what I was able to set the API
                            access through */

func main() {
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/auth/callback", handleAuthRedirect)

    fmt.Print("before")
    http.ListenAndServe(":8080", nil)
}

func handleRoot (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
	
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