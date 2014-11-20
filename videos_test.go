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

func TestListTop(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/top", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":  "1",
			"offset": "1",
			"game":   "g",
			"period": "week",
		})
		fmt.Fprint(w, `{ "_links": { "next": "https://api.twitch.tv/kraken/videos/top?game=League+of+Legends&limit=10&offset=10&period=month", "self": "https://api.twitch.tv/kraken/videos/top?game=League+of+Legends&limit=10&offset=0&period=month" }, "videos": [ { "recorded_at": "2013-03-13T09:51:31Z", "preview": "p", "description": "d", "url": "u", "title": "t", "channel": { "name": "n", "display_name": "d" }, "length": 1, "game": "g", "views": 1, "_id": "i", "_links": { "channel": "c", "self": "s" } } ] }`)
	})

	opts := &VideoRequestOptions{Game: "g", Period: "week", ListOptions: ListOptions{Limit: 1, Offset: 1}}
	got, resp, err := client.Videos.ListTop(opts)
	if err != nil {
		t.Errorf("Videos.ListTop: returned error: %v", err)
	}

	testListResponse(t, resp, nil, intPtr(10), nil)
	want := []Video{
		Video{
			ID:          stringPtr("i"),
			Preview:     stringPtr("p"),
			Description: stringPtr("d"),
			URL:         stringPtr("u"),
			Title:       stringPtr("t"),
			Game:        stringPtr("g"),
			Views:       intPtr(1),
			Length:      intPtr(1),
			RecordedAt:  stringPtr("2013-03-13T09:51:31Z"),
			Channel: &Channel{
				Name:        stringPtr("n"),
				DisplayName: stringPtr("d"),
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Videos.ListTop response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestListChannelVideos(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels/test_channel/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":      "1",
			"offset":     "1",
			"broadcasts": "true",
		})
		fmt.Fprint(w, `{ "videos": [ { "title": "t", "recorded_at": "2011-10-02T19:57:06Z", "_id": "i", "_links": { "self": "s", "owner": "o" }, "url": "u", "views": 1, "preview": "p", "length": 1, "game": "g", "description": "d" } ], "_links": { "self": "https://api.twitch.tv/kraken/channels/vanillatv/videos?limit=10&offset=0", "next": "https://api.twitch.tv/kraken/channels/vanillatv/videos?limit=10&offset=10" } }`)
	})

	opts := &VideoChannelOptions{Broadcasts: true, ListOptions: ListOptions{Limit: 1, Offset: 1}}
	got, resp, err := client.Videos.ListChannelVideos("test_channel", opts)
	if err != nil {
		t.Errorf("Videos.ListChannelVideos returned error - %v", err)
	}

	testListResponse(t, resp, nil, intPtr(10), nil)
	want := []Video{
		Video{
			ID:          stringPtr("i"),
			Preview:     stringPtr("p"),
			Description: stringPtr("d"),
			URL:         stringPtr("u"),
			Title:       stringPtr("t"),
			Game:        stringPtr("g"),
			Views:       intPtr(1),
			Length:      intPtr(1),
			RecordedAt:  stringPtr("2011-10-02T19:57:06Z"),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Videos.ListChannelVideos response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
