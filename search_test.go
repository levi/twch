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

	opts := &ListOptions{Limit: 1, Offset: 1}
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

func TestStreams(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/search/streams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":  "1",
			"offset": "1",
			"hls":    "true",
			"q":      "test",
		})
		fmt.Fprint(w, `{ "_total": 1, "streams": [ { "_id": 1, "preview": { "medium": "m", "small": "s", "large": "l", "template": "t" }, "game": "g", "channel": { "mature": null, "background": "b", "updated_at": "2013-02-15T15:22:24Z", "_id": 1, "status": "s", "logo": "l", "teams": [], "url": "u", "display_name": "d", "game": "g", "banner": "b", "name": "n", "delay": 0, "video_banner": null, "_links": { "chat": "c", "subscriptions": "s", "features": "f", "commercial": "c", "stream_key": "s", "editors": "e", "videos": "v", "self": "s", "follows": "f" }, "created_at": "2011-12-23T18:03:44Z" }, "viewers": 1, "created_at": "2014-09-12T02:03:17Z", "_links": { "self": "h" } } ], "_links": { "summary": "h", "followed": "h", "next": "https://api.twitch.tv/kraken/streams?channel=zisss%2Cvoyboy&game=Diablo+III&limit=100&offset=100", "featured": "f", "self": "https://api.twitch.tv/kraken/streams?channel=zisss%2Cvoyboy&game=Diablo+III&limit=100&offset=0" } }`)
	})

	opts := &RequestOptions{HLS: true, ListOptions: ListOptions{Limit: 1, Offset: 1}}
	got, resp, err := client.Search.Streams("test", opts)
	if err != nil {
		t.Errorf("Search.Streams: returned error: %v", err)
	}

	testListResponse(t, resp, intPtr(1), intPtr(100), nil)
	want := []Stream{
		Stream{
			ID:        intPtr(1),
			CreatedAt: stringPtr("2014-09-12T02:03:17Z"),
			Viewers:   intPtr(1),
			Game:      stringPtr("g"),
			Channel:   channelPtr(),
			Preview:   assetPtr(),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search.Streams response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestGames(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/search/games", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"q":    "test",
			"type": "suggest",
			"live": "true",
		})
		fmt.Fprint(w, `{ "games": [ { "box": { "small": "s", "large": "l", "medium": "m", "template": "t" }, "logo": { "small": "s", "large": "l", "medium": "m", "template": "t" }, "popularity": 1, "name": "n", "_id": 1, "giantbomb_id": 1 } ] }`)
	})

	got, resp, err := client.Search.Games("test", true)
	if err != nil {
		t.Errorf("Search.Games: returned error: %v", err)
	}

	testListResponse(t, resp, nil, nil, nil)
	want := []Game{
		Game{
			Name:        stringPtr("n"),
			Box:         assetPtr(),
			Logo:        assetPtr(),
			GiantbombId: intPtr(1),
			Popularity:  intPtr(1),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Search.Games response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
