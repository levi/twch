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
		fmt.Fprint(w, `{ "_links": { "self": "s", "next": "n" }, "_total": 1, "top": [ { "game": { "name": "l", "box": { "large": "l", "medium": "m", "small": "s", "template": "t" }, "logo": { "large": "l", "medium": "m", "small": "s", "template": "t" }, "_links": {}, "_id": 1, "giantbomb_id": 1 }, "viewers": 1, "channels": 1 } ] }`)
	})

	got, resp, err := client.Games.ListTop()
	if err != nil {
		t.Errorf("Games.Top: returned error: %v", err)
	}

	testResponse(t, resp)

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
