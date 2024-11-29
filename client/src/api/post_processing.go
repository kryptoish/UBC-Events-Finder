package main

import (
	"fmt"
	"strings"
	"encoding/json"
	"regexp"

	"github.com/araddon/dateparse"
)

//Lower case key terms to search for in the post caption
var KeyTerms []string = {
	"pizza",
	"donuts",
	"free snacks",
	"free burgers",
	"free pizza",
	"grab some pizza",
	"cookies",
	"muffins",
	"snacks",
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
			if strings.Contains(strings.ToLower(item.Caption), term) {
				filteredResponse.Data = append(filteredResponse.Data, item)
				break;
			}
		}
	}
	return filteredResponse
}

func relavent_info(filteredPosts MediaResponse) ProcessedResponse{
	//use reg ex to interpret dates
	/* 
		\d\d[\/|-]\d\d[\/|-](20)?(21|22|23|24|25)    DD/MM/YYYY | MM/DD/YYYY | ../../YY | ..-..-....
		[a-z]{3}\s(1st|2nd|3rd|\dth),\s

		\d{1,2}(:\d\d){0,1}([aApP][mM]){0,1} //for times 12:00AM

	*/
	var processedData ProcessedResponse

	dateExps := []string {
		`\b[A-z]{3}\s\d?(1st|2nd|3rd|\d{1,2}th)`,
		`\b\d\d[\/|-]\d\d[\/|-](20)?(21|22|23|24|25)`,
	}


	for i, postData : filteredPosts.Data {
		caption := postData.Caption
		//get time
		date, time := processDateTime(caption)

		//get location


		//get food


	}
}

/*
takes the caption and extracts the date and time using regexs
*/
func processDateTime(caption string) (string, string) {
	var time string
	var date string

	/* regexes */
	dateExps := []string {
		`\b[A-z]{3}\s\d?(1st|2nd|3rd|\d{1,2}th)`,
		`\b\d\d[\/|-]\d\d[\/|-](20)?(21|22|23|24|25)`,
	}
	postDateRegex := regexp.MustCompile(`(\d\d\d\d)-(\d\d)-(\d\d)`)
	postTimeRegex := regexp.MustCompile(`(\d\d):(\d\d):(\d\d)`)

	/* If there is no minute info, append blank minute info */
	timeReg := regexp.MustCompile(`\b(\d{1,2})(:\d\d)?\s*([aApP][mM])\b`)
	timeString := timeReg.FindStringSubmatch(caption)
	if timeString == nil {
		time = ""
	} else if timeString[2] == "" {
		minutes := ":00"
		time = fmt.Sprintf("%s%s %s", timeString[1], minutes, timeString[3])
	}

	for _, pattern := range dateExps{
		regex := regexp.MustCompile(pattern)
		if match := regex.FindString(caption); match != "" {
			println(match)
			dateTime, _ := dateparse.ParseLocal(match + " " + time)

			subsections := postDateRegex.FindStringSubmatch(dateTime.String())
			if subsections[1] == "0000" {
				date = fmt.Sprintf("%s-%s-%s", "2024", subsections[2], subsections[3])
			}
			time := postTimeRegex.FindString(dateTime.String())
			fmt.Printf("date: %s\n", date)
			fmt.Printf("time: %s\n", time)
			return date, time
		}
	}
	return date, time
}

func processLocation (caption string) string {
	locationRegex := regexp.MustCompile(`\b((Location|room):\s)([A-Za-z,\s]*[0-9]*)`)

	locations := locationRegex.FindStringSubmatch(caption)
	location := locations[2]
	return location
}