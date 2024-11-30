package main
import (
	"testing"
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

	filtered := filterData(posts)

	if len(filtered.Data) != len(expected.Data) {
		t.Errorf("Expected %d filtered posts, got %d", len(expected.Data), len(filtered.Data))
	}

	for i, post := range filtered.Data {
		if post.Caption != expected.Data[i].Caption {
			t.Errorf("Expected Caption: %s, got %s", expected.Data[i].Caption, post.Caption)
		}
	}
}

func TestProcessDateTime(t *testing.T) {
	tests := []struct {
		caption   string
		wantDate  string
		wantTime  string
	}{
		{"Event on December 1st, 2024 at 5:30PM", "2024-12-01", "17:30:00"},
		{"Date: Jan 15, Time: 12:45PM", "2024-01-15", "12:45:00"},
		{"The event is on 11/29/2024", "2024-11-29", ""},
		{"See you at noon", "", ""},
		{"No date or time provided", "", ""},
	}

	for _, test := range tests {
		gotDate, gotTime := processDateTime(test.caption)

		if gotDate != test.wantDate {
			t.Errorf("For caption '%s', expected Date: %s, got: %s", test.caption, test.wantDate, gotDate)
		}
		if gotTime != test.wantTime {
			t.Errorf("For caption '%s', expected Time: %s, got: %s", test.caption, test.wantTime, gotTime)
		}
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

	for _, test := range tests {
		gotLocation := processLocation(test.caption)

		if gotLocation != test.wantLocation {
			t.Errorf("For caption '%s', expected Location: %s, got: %s", test.caption, test.wantLocation, gotLocation)
		}
	}
}

/*
func TestRelaventInfo(t *testing.T) {
	token := "dummy_token"
	username := "test_user"
	w := httptest.NewRecorder()

	// Mock retrieveUserId and retrievePostData
	retrieveUserId = func(token string, w http.ResponseWriter) (string, error) {
		return "test_id", nil
	}
	retrievePostData = func(token, id string, w http.ResponseWriter) MediaResponse {
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
}*/