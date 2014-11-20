package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestChannels(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/search/channels", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":  "1",
			"offset": "1",
			"q":      "test",
		})
		fmt.Fprint(w, `{ "channels": [ { "updated_at": "2013-12-25T15:18:23Z", "video_banner": "b", "logo": "l", "display_name": "d", "delay": 0, "followers": 1, "status": "s", "primary_team_name": null, "views": 1, "abuse_reported": null, "game": "g", "background": null, "banner": null, "name": "n", "url": "u", "created_at": "2013-10-07T15:00:40Z", "primary_team_display_name": null, "mature": false, "profile_banner_background_color": null, "_id": 1, "profile_banner": null } ], "_total": 1, "_links": { "self": "https://api.twitch.tv/kraken/search/channels?limit=10&offset=0&q=starcraft", "next": "https://api.twitch.tv/kraken/search/channels?limit=10&offset=10&q=starcraft" } }`)
	})

	opts := &SearchOptions{ListOptions: ListOptions{Limit: 1, Offset: 1}}
	got, resp, err := client.Search.Channels("test", opts)
	if err != nil {
		t.Errorf("Search.Channels: returned error: %v", err)
	}

	testListResponse(t, resp, intPtr(1), intPtr(10), nil)
	want := []Channel{
		Channel{
			ID:          intPtr(1),
			URL:         stringPtr("u"),
			Name:        stringPtr("n"),
			DisplayName: stringPtr("d"),
			VideoBanner: stringPtr("b"),
			Logo:        stringPtr("l"),
			Delay:       intPtr(0),
			Followers:   intPtr(1),
			Status:      stringPtr("s"),
			Views:       intPtr(1),
			Game:        stringPtr("g"),
			Mature:      boolPtr(false),
			UpdatedAt:   stringPtr("2013-12-25T15:18:23Z"),
			CreatedAt:   stringPtr("2013-10-07T15:00:40Z"),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search.Channels response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
