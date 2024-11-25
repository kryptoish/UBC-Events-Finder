package main

import (
	"fmt"
	"log"

	"encoding/json"
	"github.com/go-resty/resty/v2"
)

const appID = "935845381976691"
const client_secret = "589f481d9cf782bdc1dc0d9689556218"
const redirectURI = "https://"

/**
* prompts user to click a link to get a code, then input a code
*
* Currently just prints the retrieved info and d
*/
func get_token() string {
	client := resty.New()
	client.SetDebug(true)


	link := fmt.Sprintf("https://api.instagram.com/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user_profile,user_media&response_type=code", appID, redirectURI);

	link2 := "https://www.instagram.com/oauth/authorize?enable_fb_login=0&force_authentication=1&client_id=935845381976691&redirect_uri=https://ubc-events-finder.vercel.app/&response_type=code&scope=instagram_business_basic"

	var code string;

	fmt.Printf("Click this link:\n%s\nOr this link:\n%s\n", link, link2)
	fmt.Print("Then, input the code: ")

	fmt.Scan(&code)
	
	response, err := client.R().
		SetQueryParams(map[string]string{
			"client_id": appID,
			"client_secret": client_secret,
			"grant_type": "authorization_code",
			"redirect_uri":  redirectURI,
			"code": code,
		}).
		Post("https://api.instagram.com/oauth/access_token")
	if err != nil {
		log.Fatalf("error getting access_token:    %v", err)
	}

	fmt.Printf("response body\n %s \n", response.Body())

	var userData map[string] interface{}
	if err := json.Unmarshal(response.Body(), &userData); err != nil {
		log.Fatalf("error parsing user ID response.   %v:", err)
	}

	return "a"//response.access_token

/*
	userID, ok := userData["id"].(string)
	if !ok {
		log.Fatalf("unable to get user ID from response")
	}

	mediaResponse, err := client.R().
		SetQueryParams(map[string]string{
			"fields":	   "id,caption,media_url,permalink",
			"access_token": accessToken,
		}).
		Get(fmt.Sprintf("https://graph.instagram.com/v16.0/%s/media", userID))
	
	if err != nil {
		log.Fatalf("Error getting media:   %v", err)
	}

	var mediaData MediaResponse
	if err := json.Unmarshal(mediaResponse.Body(), &mediaData); err != nil {
		log.Fatalf("Error parsing media response:   %v", err)
	}


	for _, post := range mediaData.Data {
		fmt.Printf("Caption: %s\n", post.Caption)
		fmt.Printf("Media URL: %s\n", post.MediaURL)
		fmt.Printf("Permalink: %s\n", post.Permalink)
		fmt.Printf("------\n")
	}*/
}
