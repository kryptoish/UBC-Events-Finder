package main

import (
	"fmt"
	"strings"
	"regexp"

	"github.com/araddon/dateparse"
)

//Lower case key terms to search for in the post caption
var KeyTerms = []string{
	"pizza",
	"donuts",
	"free snacks",
	"free burgers",
	"free pizza",
	"grab some pizza",
	"cookies",
	"muffins",
	"snacks",
	"BBQ",
	"hotdogs",
	"burgers",
}

var FoodTerms = []string{
	"Pizza",
	"Donuts",
	"Cookies",
	"Muffins",
	"Snacks",
	"Red bull",
	"Candy",
	"Hotdogs",
	"Burger",
}

/**
* 
* 
* 
*/
func filter_data(posts MediaResponse) MediaResponse {
	var filteredResponse MediaResponse

	for _, item := range posts.Data {
		for _, term := range KeyTerms {
			if strings.Contains(strings.ToLower(item.Caption), strings.ToLower(term)) {
				filteredResponse.Data = append(filteredResponse.Data, item)
				break;
			}
		}
	}
	return filteredResponse
}

func relavent_info(filteredPosts MediaResponse, username string) ProcessedResponse {
	var processedData ProcessedResponse

	processedData.Data = make([]struct {
		ID        string `json:"id"`
		Caption   string `json:"caption"`
		MediaURL  string `json:"media_url"`
		Permalink string `json:"permalink"`
		Username  string `json:"username"`
		Food      string `json:"food"`
		Date      string `json:"date"`
		Time      string `json:"time"`
		Location  string `json:"location"`
	}, len(filteredPosts.Data))

	for i, postData := range filteredPosts.Data {
		caption := postData.Caption
		//get time
		date, time := processDateTime(caption)

		//get location
		location := processLocation(caption)

		//get food
		var foods []string
		for _, food := range FoodTerms {
			if strings.Contains(strings.ToLower(caption), strings.ToLower(food)) {
				foods = append(foods, food)
			}
		}
		food := strings.Join(foods, ", ")
		
		processedData.Data[i].ID        = postData.ID;
		processedData.Data[i].Caption   = postData.Caption;
		processedData.Data[i].MediaURL  = postData.MediaURL;
		processedData.Data[i].Permalink = postData.Permalink;
		processedData.Data[i].Username  = username;
		processedData.Data[i].Food      = food;
		processedData.Data[i].Date      = date;
		processedData.Data[i].Time      = time;
		processedData.Data[i].Location  = location;
	}
	return processedData
}

/*
takes the caption and extracts the date and time using regexs
*/
func processDateTime(caption string) (string, string) {
	var time string
	var tempTime string
	var date string

	/* regexes */
	dateExps := []string {
		`\b[A-z]{3}\s\d?(1st|2nd|3rd|\d{1,2}th)`,
		`\d\d[\/|-]\d\d[\/|-](20)?(21|22|23|24|25)`,
		`((?i)date):?\s+([a-zA-Z0-9,]+ [a-zA-Z0-9]+)`,
	}
	postDateRegex := regexp.MustCompile(`(\d\d\d\d)-(\d\d)-(\d\d)`)
	postTimeRegex := regexp.MustCompile(`(\d\d):(\d\d):(\d\d)`)

	/* If there is no minute info, append blank minute info */
	timeReg := regexp.MustCompile(`(\d{1,2})(:\d\d)?\s*([aApP][mM])`)
	timeString := timeReg.FindStringSubmatch(caption)
	if timeString == nil {
		time = ""
		tempTime = "00:00"
	} else if timeString[2] == "" {
		minutes := ":00"
		time = fmt.Sprintf("%s%s %s", timeString[1], minutes, timeString[3])
		tempTime = time
	}

	for _, pattern := range dateExps{
		regex := regexp.MustCompile(pattern)
		if match := regex.FindStringSubmatch(caption); match != nil {
			println(match[0] + " " + tempTime)
			dateTime, _ := dateparse.ParseLocal(match[0] + " " + tempTime)
			if strings.Contains(strings.ToLower(match[0]), strings.ToLower("date")) {
				dateTime, _ = dateparse.ParseLocal(match[2] + " " + tempTime)
			}
				
			subsections := postDateRegex.FindStringSubmatch(dateTime.String())
			println(dateTime.String())
			if subsections[1] == "0000" {
				date = fmt.Sprintf("%s-%s-%s", "2024", subsections[2], subsections[3])
			} 

			tempTime := postTimeRegex.FindString(dateTime.String())
			fmt.Printf("date: %s\n", date)
			fmt.Printf("time: %s\n", tempTime)
			return date, tempTime
		}
	}
	return date, time
}

func processLocation (caption string) string {
	locationRegex := regexp.MustCompile(`((?i)(location|room):?\s+)([A-Za-z, ]*[0-9]*)`)
	var location string

	locations := locationRegex.FindStringSubmatch(caption)
	if locations != nil {
		location = locations[3]
	} else {
		location = ""
	}
	return location
}

func mergedResponses (responses []ProcessedResponse) ProcessedResponse {
	var merged ProcessedResponse

	for _, response := range responses {
		merged.Data = append(merged.Data, response.Data...)
	}

	return merged
}