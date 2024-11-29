package main

import (
	"fmt"
	"strings"
	"regexp"
	"time"

	//"github.com/araddon/dateparse"
	"github.com/markusmobius/go-dateparser"
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
	cfg := &dateparser.Configuration{
		CurrentTime: time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC),
	}

	var time string
	var tempTime string
	var date string

	/* regexes */
	dateExps := []string {
		`\b[A-z]{3}\s\d?(1st|2nd|3rd|\d{1,2}th)`,
		`\d\d[\/|-]\d\d[\/|-](20)?(21|22|23|24|25)`,
		`((?i)date):?\s+([a-zA-Z0-9,]+ [a-zA-Z0-9]+(,?\s\d\d\d\d)?)`,
	}
	postDateRegex := regexp.MustCompile(`(\d\d\d\d)-(\d\d)-(\d\d)`)
	postTimeRegex := regexp.MustCompile(`(\d\d):(\d\d):(\d\d)`)

	/* If there is no minute info, append blank minute info */
	timeReg := regexp.MustCompile(`((\d{1,2})(:\d\d)?\s*([aApP][mM])|(\d{1,2})(:\d\d)\s*([aApP][mM])?)`)//-((\d{1,2})(:\d\d)?\s*([aApP][mM])|(\d{1,2})(:\d\d)\s*[aApP][mM])?`)
	timeString := timeReg.FindStringSubmatch(caption)
	fmt.Printf("timeString: %s\n", timeString[0])
	if timeString == nil {
		time = ""
		tempTime = "00:00"
	} else if !strings.Contains(timeString[0], ":") {
		minutes := ":00"
		tempTime = fmt.Sprintf("%s%s %s", timeString[2], minutes, timeString[4])
	} else {
		tempTime = timeString[0]
	}
	time = tempTime

	//potential way to avoid the regexs but performace is subpar
	//_, dates, _ := dateparser.Search(cfg, caption)
	//println(dates[0].Date.Time.String())

	for _, pattern := range dateExps{
		regex := regexp.MustCompile(pattern)
		if match := regex.FindStringSubmatch(caption); match != nil {
			yearExp := regexp.MustCompile(`\d\d\d\d`)			
			if year := yearExp.FindString(match[0]); year == "" {
				match[0] = match[0] + " 2024"
				match[2] = match[2] + " 2024"
			}


			dateTime, _ := dateparser.Parse(cfg, match[0] + " " + time)
			if strings.Contains(strings.ToLower(match[0]), strings.ToLower("date")) {
				dateTime, _ = dateparser.Parse(cfg, match[2] + " " + time)
			}
				
			subsections := postDateRegex.FindStringSubmatch(dateTime.Time.String())
			println(dateTime.Time.String())
			if subsections[1] != "2024" {
				date = fmt.Sprintf("%s-%s-%s", "2024", subsections[2], subsections[3])
			} else {
				date = subsections[0]
			}

			time := postTimeRegex.FindString(dateTime.Time.String())

			return date, time
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