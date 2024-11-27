package main

import (
	"log"
	"os"
	"fmt"

	"net/url"
	"net/http"
	"encoding/json"
)

const redirect_uri = "https://ubc-events-finder.vercel.app/auth/callback/"


/**
* prompts user to click a link to get a code, then input a code
*
* Currently just prints the retrieved info and d
*/
func get_token(code string) string {
	fmt.Printf("get token from code: %s\n", code)
	appID := os.Getenv("APP_ID")
	client_secret := os.Getenv("CLIENT_SECRET")


	//link := fmt.Sprintf("https://api.instagram.com/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user_profile,user_media&response_type=code", appID, redirectURI);
	short_form_data := url.Values{
		"client_id":     {appID},
		"client_secret": {client_secret},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {redirect_uri},
		"code":          {code},
	}
	short_term_auth_url := "https://api.instagram.com/oauth/access_token" 
	short_term_token := exchange_token(short_term_auth_url, short_form_data)


	
	long_form_data := url.Values{
		"client_secret": {client_secret},
		"grant_type":    {"ig_exchange_token"},
		"access_token":  {short_term_token},
	}
	long_term_auth_url := "https://graph.instagram.com/oauth/access_token" 
	long_term_token := exchange_token(long_term_auth_url, long_form_data)
	
	return long_term_token
}

func exchange_token(url string, form_data url.Values) string {
	response, err := http.PostForm(url, form_data)
	if err != nil {
		fmt.Printf("error getting access_token:    %v", err)
		log.Fatalf("error getting access_token:    %v", err)
	}
	defer response.Body.Close()
	
	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		fmt.Printf("Error parsing response: %v", err)
		log.Fatalf("Error parsing response: %v", err)
	}

	token, ok := data["access_token"].(string)
	if !ok {
		fmt.Printf("Unable to get token from response")
		log.Fatalf("Unable to get token from response")
	}

	return token
}
