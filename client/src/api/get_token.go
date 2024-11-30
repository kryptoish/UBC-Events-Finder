package main

import (
	"os"
	"errors"

	"net/url"
	"net/http"
	"encoding/json"
)

const redirectUri = "https://ubc-events-finder.vercel.app/auth/callback/"


/* When passed a valid, unused authorization code 
	attempts to return a long term access token to 
	the instagram api */
func getToken(code string) (string, error) {
	appID := os.Getenv("APP_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	shortFormData := url.Values{
		"client_id":     {appID},
		"client_secret": {clientSecret},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {redirectUri},
		"code":          {code},
	}
	shortTermAuthUrl := "https://api.instagram.com/oauth/access_token" 
	shortTermToken, err := exchangeToken(shortTermAuthUrl, shortFormData)
	if err != nil {
		return "", err
	}
	
	longFormData := url.Values{
		"client_secret": {clientSecret},
		"grant_type":    {"ig_exchange_token"},
		"access_token":  {shortTermToken},
	}
	longTermAuthUrl := "https://graph.instagram.com/oauth/access_token" 
	longTermToken, err := exchangeToken(longTermAuthUrl, longFormData)
	if err != nil {
		return "", err
	}
	
	return longTermToken, nil
}

/* Takes URL values and sends a POST request to the 
	URL endpoint provided. Forwards any erros on */
func exchangeToken(url string, formData url.Values) (string, error) {
	response, err := http.PostForm(url, formData)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	
	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return "", err
	}

	token, ok := data["access_token"].(string)
	if !ok {
		return "", errors.New("Unable to get token from response\n")
	}

	return token, nil
}
