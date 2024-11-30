package main

import (
	"fmt"
	"io"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

/* From an accessToken, returns the userID and username 
	Requires a response writer to pass http errors back */
func retrieveUserId(accessToken string, w http.ResponseWriter) (string, string) {
	userIdUrl := fmt.Sprintf("https://graph.instagram.com/v21.0/me?fields=id,username&access_token=%s", accessToken)
	response, err := http.Get(userIdUrl)
	if err != nil {
		errorMsg := fmt.Sprintf("Error getting user ID: %v\n", err)
        http.Error(w, errorMsg, http.StatusBadRequest)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		errorMsg := fmt.Sprintf("Error response getting user ID: %s\n", body)
        http.Error(w, errorMsg, http.StatusBadRequest)
	}

	var responseDecoded map[string]interface{}	
	if err := json.NewDecoder(response.Body).Decode(&responseDecoded); err != nil {
		errorMsg := fmt.Sprintf("Error decoding response: %v\n", err)
        http.Error(w, errorMsg, http.StatusBadRequest)
	}

	userId, ok := responseDecoded["id"].(string)
	if !ok {
		errorMsg := fmt.Sprintf("Error parsing user ID response: %v\n", err)
        http.Error(w, errorMsg, http.StatusBadRequest)
	}	
	username, _ := responseDecoded["username"].(string)

	return userId, username
}


/* From an accessToken and a userID, returns a MediaResponse 
	containing all the post information from the user.
	Requires a response writer to pass http errors back */
func retrievePostData(accessToken, userID string, w http.ResponseWriter) MediaResponse {
	/* get user ID */
	mediaURL := fmt.Sprintf("https://graph.instagram.com/v21.0/%s/media?fields=id,caption,media_url,permalink&access_token=%s", userID, accessToken)
	mediaResponse, err := http.Get(mediaURL)
	if err != nil {
		errorMsg := fmt.Sprintf("Error getting media: %v\n", err)
        http.Error(w, errorMsg, http.StatusBadRequest)
	}
	defer mediaResponse.Body.Close()

	if mediaResponse.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(mediaResponse.Body)
		errorMsg := fmt.Sprintf("Error response getting media: %s\n", body)
        http.Error(w, errorMsg, http.StatusBadRequest)
	}

	mediaBody, err := ioutil.ReadAll(mediaResponse.Body)
	if err != nil {
		errorMsg := fmt.Sprintf("Error reading media response body: %v\n", err)
        http.Error(w, errorMsg, http.StatusBadRequest)
	}

	var mediaData MediaResponse
	if err := json.Unmarshal(mediaBody, &mediaData); err != nil {
		errorMsg := fmt.Sprintf("Error parsing media response: %v\n", err)
        http.Error(w, errorMsg, http.StatusBadRequest)
	}

	return mediaData
}