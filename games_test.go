package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTopGames(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/games/top", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":  "1",
			"offset": "1",
			"hls":    "true",
		})
		fmt.Fprint(w, `{ "_links": { "self": "s", "next": "https://api.twitch.tv/kraken/games/top?limit=10&offset=10" }, "_total": 1, "top": [ { "game": { "name": "l", "box": { "large": "l", "medium": "m", "small": "s", "template": "t" }, "logo": { "large": "l", "medium": "m", "small": "s", "template": "t" }, "_links": {}, "_id": 1, "giantbomb_id": 1 }, "viewers": 1, "channels": 1 } ] }`)
	})

	opts := &RequestOptions{ListOptions: ListOptions{Limit: 1, Offset: 1}, HLS: true}
	got, resp, err := client.Games.ListTop(opts)
	if err != nil {
		t.Errorf("Games.Top: returned error: %v", err)
	}

	testListResponse(t, resp, intPtr(1), intPtr(10), nil)

	want := []Game{
		Game{
			Name:        stringPtr("l"),
			Box:         &Asset{Large: stringPtr("l"), Medium: stringPtr("m"), Small: stringPtr("s"), Template: stringPtr("t")},
			Logo:        &Asset{Large: stringPtr("l"), Medium: stringPtr("m"), Small: stringPtr("s"), Template: stringPtr("t")},
			GiantbombId: intPtr(1),
			Viewers:     intPtr(1),
			Channels:    intPtr(1),
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Games.Top: response was wrong: %+v", got)
	}
}
