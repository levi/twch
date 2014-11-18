package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetChannel(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/channels/test_user1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "name": "n", "game": "g", "created_at": "2011-02-24T01:38:43Z", "title": "t", "updated_at": "2012-06-18T05:22:53Z", "banner": "b", "video_banner": "v", "background": "b", "_links": { "self": "h", "chat": "h", "videos": "h", "commercial": "h", "follows":"h", "stream_key":"h", "features":"h", "subscriptions":"h", "editors":"h" }, "logo": "l", "_id": 1, "mature": true, "url": "u", "display_name": "d" }`)
	})
	want := &Channel{
		Name:        stringPtr("n"),
		Game:        stringPtr("g"),
		Banner:      stringPtr("b"),
		Title:       stringPtr("t"),
		VideoBanner: stringPtr("v"),
		Background:  stringPtr("b"),
		Logo:        stringPtr("l"),
		ID:          intPtr(1),
		Mature:      boolPtr(true),
		DisplayName: stringPtr("d"),
		URL:         stringPtr("u"),
		CreatedAt:   stringPtr("2011-02-24T01:38:43Z"),
		UpdatedAt:   stringPtr("2012-06-18T05:22:53Z"),
	}
	got, _, err := client.Channels.GetChannel("test_user1")
	if err != nil {
		t.Errorf("Channels.GetChannel: request returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Channels.GetChannel response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
