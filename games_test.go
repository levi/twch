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
		if r.Method != "GET" {
			t.Errorf("Top Games: method = %s, want %v", r.Method, "GET")
		}
		fmt.Fprint(w, `{ "_links": { "self": "s", "next": "n" }, "_total": 1, "top": [ { "game": { "name": "l", "box": { "large": "l", "medium": "m", "small": "s", "template": "t" }, "logo": { "large": "l", "medium": "m", "small": "s", "template": "t" }, "_links": {}, "_id": 1, "giantbomb_id": 1 }, "viewers": 1, "channels": 1 } ] }`)
	})

	got, _, err := client.Games.Top()
	if err != nil {
		t.Errorf("Games.Top: returned error: %v", err)
	}

	want := []Game{
		Game{Name: "l", Box: &Asset{Large: "l", Medium: "m", Small: "s", Template: "t"}, Logo: &Asset{Large: "l", Medium: "m", Small: "s", Template: "t"}, GiantbombId: 1, Viewers: 1, Channels: 1},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Games.Top: response was wrong: %+v", got)
	}
}
