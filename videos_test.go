package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetVideo(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "recorded_at": "2012-08-09T20:49:47Z", "title": "t", "url": "u", "_id": "i", "_links": { "self": "s", "owner": "o" }, "views": 1, "description": "d", "length": 1, "game": "g", "preview": "p" }`)
	})
	want := &Video{
		ID:          stringPtr("i"),
		URL:         stringPtr("u"),
		Title:       stringPtr("t"),
		Views:       intPtr(1),
		Description: stringPtr("d"),
		Length:      intPtr(1),
		Game:        stringPtr("g"),
		Preview:     stringPtr("p"),
		RecordedAt:  stringPtr("2012-08-09T20:49:47Z"),
	}
	got, _, err := client.Videos.GetVideo(1)
	if err != nil {
		t.Errorf("Videos.GetVideo: request returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Videos.GetVideo response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
