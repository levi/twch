package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListTeams(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "_links": { "next": "https://api.twitch.tv/kraken/teams?limit=25&offset=25", "self": "https://api.twitch.tv/kraken/teams?limit=25&offset=0" }, "teams": [ { "info": "i", "_links": { "self": "s" }, "background": "b", "banner": "b", "name": "n", "_id": 1, "updated_at": "2012-11-14T01:30:00Z", "display_name": "d", "created_at": "2011-10-11T22:49:05Z", "logo": "l" } ] }`)
	})

	got, resp, err := client.Teams.ListTeams()
	if err != nil {
		t.Errorf("Teams.ListTeams: returned error: %v", err)
	}

	testListResponse(t, resp, nil, intPtr(25), nil)
	want := []Team{
		Team{
			ID:          intPtr(1),
			Name:        stringPtr("n"),
			DisplayName: stringPtr("d"),
			Banner:      stringPtr("b"),
			Background:  stringPtr("b"),
			Info:        stringPtr("i"),
			Logo:        stringPtr("l"),
			CreatedAt:   stringPtr("2011-10-11T22:49:05Z"),
			UpdatedAt:   stringPtr("2012-11-14T01:30:00Z"),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Teams.ListTeams response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestGetTeam(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/teams/test_team", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "_id": 1, "created_at": "2011-10-11T23:59:43Z", "info": "i", "updated_at": "2012-01-15T19:43:40Z", "background": "b", "banner": "b", "logo": "l", "_links": { "self": "s" }, "display_name": "d", "name": "n" }`)
	})
	want := &Team{
		ID:          intPtr(1),
		Info:        stringPtr("i"),
		Background:  stringPtr("b"),
		Banner:      stringPtr("b"),
		Logo:        stringPtr("l"),
		DisplayName: stringPtr("d"),
		Name:        stringPtr("n"),
		CreatedAt:   stringPtr("2011-10-11T23:59:43Z"),
		UpdatedAt:   stringPtr("2012-01-15T19:43:40Z"),
	}
	got, _, err := client.Teams.GetTeam("test_team")
	if err != nil {
		t.Errorf("Teams.GetTeam: request returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Teams.GetTeam response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
