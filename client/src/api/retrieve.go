package main

import (
	"fmt"
	"log"
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
func retrieve_post_data(access_token, userID string) MediaResponse {
	/* get user ID */
	mediaURL := fmt.Sprintf("https://graph.instagram.com/v21.0/%s/media?fields=id,caption,media_url,permalink&access_token=%s", userID, access_token)
	mediaResponse, err := http.Get(mediaURL)
	if err != nil {
		log.Fatalf("Error getting media: %v", err)
	}
	defer mediaResponse.Body.Close()

	if mediaResponse.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(mediaResponse.Body)
		log.Fatalf("Error response getting media: %s", body)
	}

	mediaBody, err := ioutil.ReadAll(mediaResponse.Body)
	if err != nil {
		log.Fatalf("Error reading media response body: %v", err)
	}

	var mediaData MediaResponse
	if err := json.Unmarshal(mediaBody, &mediaData); err != nil {
		log.Fatalf("Error parsing media response: %v", err)
	}

	for _, post := range mediaData.Data {
		fmt.Printf("Caption: %s\n", post.Caption)
		fmt.Printf("Media URL: %s\n", post.MediaURL)
		fmt.Printf("Permalink: %s\n", post.Permalink)
	}

	return mediaData
}

func retrieve_user_id(access_token string) string {
	user_id_url := fmt.Sprintf("https://graph.instagram.com/v21.0/me?fields=id,username&access_token=%s", access_token)
	response, err := http.Get(user_id_url)
	if err != nil {
		log.Fatalf("Error getting user ID: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		log.Fatalf("Error response getting user ID: %s", body)
	}

	var response_decoded map[string]interface{}	
	if err := json.NewDecoder(response.Body).Decode(&response_decoded); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	user_id, ok := response_decoded["id"].(string)
	if !ok {
		log.Fatalf("Error parsing user ID response: %v", err)
	}	

	return user_id
}