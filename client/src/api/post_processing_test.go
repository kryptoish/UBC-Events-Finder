package main

import (
	"testing"
	"strings"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestFilterData(t *testing.T) {
	posts := MediaResponse{
		Data: []struct {
			ID        string `json:"id"`
			Caption   string `json:"caption"`
			MediaURL  string `json:"media_url"`
			Permalink string `json:"permalink"`
		}{
			{ID: "1", Caption: "Free pizza at the event!", MediaURL: "url1", Permalink: "link1"},
			{ID: "2", Caption: "No free food here.", MediaURL: "url2", Permalink: "link2"},
			{ID: "3", Caption: "Come grab some free burgers and cookies!", MediaURL: "url3", Permalink: "link3"},
			{ID: "4", Caption: "Tickets are $10", MediaURL: "url4", Permalink: "link4"},
		},
	}

	expected := MediaResponse{
		Data: []struct {
			ID        string `json:"id"`
			Caption   string `json:"caption"`
			MediaURL  string `json:"media_url"`
			Permalink string `json:"permalink"`
		}{
			{ID: "1", Caption: "Free pizza at the event!", MediaURL: "url1", Permalink: "link1"},
			{ID: "3", Caption: "Come grab some free burgers and cookies!", MediaURL: "url3", Permalink: "link3"},
		},
	}
	a := assert.New(t)

	filtered := filterData(posts)
	a.Equal(len(filtered.Data), len(expected.Data))

	for i, post := range filtered.Data {
		a.Equal(post.Caption, expected.Data[i].Caption)
	}
}

func TestProcessDateTime(t *testing.T) {
	tests := []struct {
		caption   string
		wantDate  string
		wantTime  string
	}{
		{"Event on December 1st, 2024 at 5:30PM", "2024-12-01", "17:30:00"},
		{"Event on December 1st, 2024 at 5PM", "2024-12-01", "17:00:00"},
		{"Event on December 1st at 5PM", "2024-12-01", "17:00:00"},
		{"Date: Jan 15, Time: 12:45PM", "2024-01-15", "12:45:00"},
		{"Big things on January 15, Time: 12:45PM", "2024-01-15", "12:45:00"},
		{"Big things on January 15, 2026 Time: 12:45PM", "2026-01-15", "12:45:00"},
		{"The event is on 11/29/2024", "2024-11-29", ""},
		{"See you at noon", "", ""},
		{"No date or time provided", "", ""},
		{"March 12 nothin else", "2024-03-12", ""},
	}

	a := assert.New(t)
	for _, test := range tests {
		gotDate, gotTime := processDateTime(test.caption)
		a.Equal(gotDate, test.wantDate)
		a.Equal(gotTime, test.wantTime)
	}
}

func TestProcessLocation(t *testing.T) {
	tests := []struct {
		caption       string
		wantLocation  string
	}{
		{"Location: Room 123", "Room 123"},
		{"Event in Building 456", "Building 456"},
		{"Meet at the cafe", ""},
		{"No location provided", "provided"},
	}

	a := assert.New(t)
	for _, test := range tests {
		gotLocation := processLocation(test.caption)
		a.Equal(gotLocation, test.wantLocation)
	}
}


func TestRelaventInfo(t *testing.T) {
	token := "dummy_token"
	username := "test_user"
	w := httptest.NewRecorder()
	
	// Mock retrieveUserId and retrievePostData
	RetrieveUserId = func(token string, w http.ResponseWriter) (string, string) {
		return "test_id", ""
	}
	RetrievePostData = func(token, id string, w http.ResponseWriter) MediaResponse {
		return MediaResponse{
			Data: []struct {
				ID        string `json:"id"`
				Caption   string `json:"caption"`
				MediaURL  string `json:"media_url"`
				Permalink string `json:"permalink"`
			}{
				{ID: "1", Caption: "Free pizza and burgers at noon!", MediaURL: "url1", Permalink: "link1"},
			},
		}
	}

	processed := relaventInfo(token, username, w)

	if len(processed.Data) != 1 {
		t.Fatalf("Expected 1 processed post, got %d", len(processed.Data))
	}

	post := processed.Data[0]

	if post.Username != username {
		t.Errorf("Expected Username: %s, got %s", username, post.Username)
	}

	if !strings.Contains(post.Food, "Pizza") || !strings.Contains(post.Food, "Burgers") {
		t.Errorf("Expected Food: Pizza, Burgers; got: %s", post.Food)
	}

	if post.Location != "" {
		t.Errorf("Expected Location: empty, got: %s", post.Location)
	}
}

func TestMergedResponses(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		inputResponses []ProcessedResponse
		expected       ProcessedResponse
	}{
		{
			name: "Merge two responses",
			inputResponses: []ProcessedResponse{
				{
					Data: []struct {
						ID        string `json:"id"`
						Caption   string `json:"caption"`
						MediaURL  string `json:"media_url"`
						Permalink string `json:"permalink"`
						Username  string `json:"username"`
						Food      string `json:"food"`
						Date      string `json:"date"`
						Time      string `json:"time"`
						Location  string `json:"location"`
					}{
						{
							ID:        "1",
							Caption:   "Caption 1",
							MediaURL:  "https://example.com/media1",
							Permalink: "https://example.com/post1",
							Username:  "user1",
							Food:      "Pizza",
							Date:      "2024-11-29",
							Time:      "12:00 PM",
							Location:  "Location 1",
						},
					},
				},
				{
					Data: []struct {
						ID        string `json:"id"`
						Caption   string `json:"caption"`
						MediaURL  string `json:"media_url"`
						Permalink string `json:"permalink"`
						Username  string `json:"username"`
						Food      string `json:"food"`
						Date      string `json:"date"`
						Time      string `json:"time"`
						Location  string `json:"location"`
					}{
						{
							ID:        "2",
							Caption:   "Caption 2",
							MediaURL:  "https://example.com/media2",
							Permalink: "https://example.com/post2",
							Username:  "user2",
							Food:      "Burger",
							Date:      "2024-11-30",
							Time:      "1:00 PM",
							Location:  "Location 2",
						},
					},
				},
			},
			expected: ProcessedResponse{
				Data: []struct {
					ID        string `json:"id"`
					Caption   string `json:"caption"`
					MediaURL  string `json:"media_url"`
					Permalink string `json:"permalink"`
					Username  string `json:"username"`
					Food      string `json:"food"`
					Date      string `json:"date"`
					Time      string `json:"time"`
					Location  string `json:"location"`
				}{
					{
						ID:        "1",
						Caption:   "Caption 1",
						MediaURL:  "https://example.com/media1",
						Permalink: "https://example.com/post1",
						Username:  "user1",
						Food:      "Pizza",
						Date:      "2024-11-29",
						Time:      "12:00 PM",
						Location:  "Location 1",
					},
					{
						ID:        "2",
						Caption:   "Caption 2",
						MediaURL:  "https://example.com/media2",
						Permalink: "https://example.com/post2",
						Username:  "user2",
						Food:      "Burger",
						Date:      "2024-11-30",
						Time:      "1:00 PM",
						Location:  "Location 2",
					},
				},
			},
		},
		{
			name:           "Empty input",
			inputResponses: []ProcessedResponse{},
			expected:       ProcessedResponse{Data: nil},
		},
		{
			name: "Single response",
			inputResponses: []ProcessedResponse{
				{
					Data: []struct {
						ID        string `json:"id"`
						Caption   string `json:"caption"`
						MediaURL  string `json:"media_url"`
						Permalink string `json:"permalink"`
						Username  string `json:"username"`
						Food      string `json:"food"`
						Date      string `json:"date"`
						Time      string `json:"time"`
						Location  string `json:"location"`
					}{
						{
							ID:        "3",
							Caption:   "Caption 3",
							MediaURL:  "https://example.com/media3",
							Permalink: "https://example.com/post3",
							Username:  "user3",
							Food:      "Sushi",
							Date:      "2024-11-28",
							Time:      "6:00 PM",
							Location:  "Location 3",
						},
					},
				},
			},
			expected: ProcessedResponse{
				Data: []struct {
					ID        string `json:"id"`
					Caption   string `json:"caption"`
					MediaURL  string `json:"media_url"`
					Permalink string `json:"permalink"`
					Username  string `json:"username"`
					Food      string `json:"food"`
					Date      string `json:"date"`
					Time      string `json:"time"`
					Location  string `json:"location"`
				}{
					{
						ID:        "3",
						Caption:   "Caption 3",
						MediaURL:  "https://example.com/media3",
						Permalink: "https://example.com/post3",
						Username:  "user3",
						Food:      "Sushi",
						Date:      "2024-11-28",
						Time:      "6:00 PM",
						Location:  "Location 3",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := mergedResponses(test.inputResponses)
			assert.Equal(t, test.expected, result)
		})
	}
}
