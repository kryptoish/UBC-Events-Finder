package main

import (
	"fmt"
	"log"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
* Takes 2 strings, the token for the account, and the username
*
* Currently just prints the retrieved info and doesn't return anything
*/
func retrieve_post_data(accessToken, username string) MediaResponse {
	/* get user ID */
	fmt.Println("Start of retrieve")

	userIDURL := fmt.Sprintf("https://graph.instagram.com/v21.0/me?fields=id,username&access_token=%s", accessToken)
	userIDResponse, err := http.Get(userIDURL)
	if err != nil {
		log.Fatalf("Error getting user ID: %v", err)
	}
	defer userIDResponse.Body.Close()

	if userIDResponse.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(userIDResponse.Body)
		log.Fatalf("Error response getting user ID: %s", body)
	}

	userIDBody, err := ioutil.ReadAll(userIDResponse.Body)
	if err != nil {
		log.Fatalf("Error reading user ID response body: %v", err)
	}

	var userData map[string]interface{}
	if err := json.Unmarshal(userIDBody, &userData); err != nil {
		log.Fatalf("Error parsing user ID response: %v", err)
	}

	userID, ok := userData["id"].(string)
	if !ok {
		log.Fatalf("Unable to get user ID from response")
	}

	/* get media data */
	mediaURL := fmt.Sprintf("https://graph.instagram.com/v21.0/%s/media?fields=id,caption,media_url,permalink&access_token=%s", userID, accessToken)
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
		fmt.Printf("------\n")
	}

	return mediaData
}