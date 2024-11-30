package main

import (
	"fmt"
	"io"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
* Takes 2 strings, the token for the account, and the username
*
* Currently just prints the retrieved info and doesn't return anything
*/
func retrieve_post_data(access_token, userID string, w http.ResponseWriter) MediaResponse {
	/* get user ID */
	mediaURL := fmt.Sprintf("https://graph.instagram.com/v21.0/%s/media?fields=id,caption,media_url,permalink&access_token=%s", userID, access_token)
	mediaResponse, err := http.Get(mediaURL)
	if err != nil {
		error_msg := fmt.Sprintf("Error getting media: %v\n", err)
        http.Error(w, error_msg, http.StatusBadRequest)
		//log.Fatalf("Error getting media: %v", err)
	}
	defer mediaResponse.Body.Close()

	if mediaResponse.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(mediaResponse.Body)
		error_msg := fmt.Sprintf("Error response getting media: %s\n", body)
        http.Error(w, error_msg, http.StatusBadRequest)
		//log.Fatalf("Error response getting media: %s", body)
	}

	mediaBody, err := ioutil.ReadAll(mediaResponse.Body)
	if err != nil {
		error_msg := fmt.Sprintf("Error reading media response body: %v\n", err)
        http.Error(w, error_msg, http.StatusBadRequest)
		//log.Fatalf("Error reading media response body: %v", err)
	}

	var mediaData MediaResponse
	if err := json.Unmarshal(mediaBody, &mediaData); err != nil {
		error_msg := fmt.Sprintf("Error parsing media response: %v\n", err)
        http.Error(w, error_msg, http.StatusBadRequest)
		//log.Fatalf("Error parsing media response: %v", err)
	}

	return mediaData
}

func retrieve_user_id(access_token string, w http.ResponseWriter) (string, string) {
	user_id_url := fmt.Sprintf("https://graph.instagram.com/v21.0/me?fields=id,username&access_token=%s", access_token)
	response, err := http.Get(user_id_url)
	if err != nil {
		error_msg := fmt.Sprintf("Error getting user ID: %v\n", err)
        http.Error(w, error_msg, http.StatusBadRequest)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		error_msg := fmt.Sprintf("Error response getting user ID: %s\n", body)
        http.Error(w, error_msg, http.StatusBadRequest)
		//log.Fatalf("Error response getting user ID: %s", body)
	}

	var response_decoded map[string]interface{}	
	if err := json.NewDecoder(response.Body).Decode(&response_decoded); err != nil {
		error_msg := fmt.Sprintf("Error decoding response: %v\n", err)
        http.Error(w, error_msg, http.StatusBadRequest)
		//log.Fatalf("Error decoding response: %v", err)
	}

	user_id, ok := response_decoded["id"].(string)
	if !ok {
		error_msg := fmt.Sprintf("Error parsing user ID response: %v\n", err)
        http.Error(w, error_msg, http.StatusBadRequest)
		//log.Fatalf("Error parsing user ID response: %v", err)
	}	
	username, _ := response_decoded["username"].(string)

	return user_id, username
}