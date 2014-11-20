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
		t.Errorf("Teams.ListTeams: response was wrong: %+v", got)
	}
}
