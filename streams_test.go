package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStreamGetSummary(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/streams/summary", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "viewers": 1, "_links": { "self": "h" }, "channels": 1 }`)
	})
	want := &StreamSummary{Viewers: 1, Channels: 1}
	got, _, err := client.Streams.GetSummary()
	if err != nil {
		t.Errorf("Streams.Summary returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Streams.Summary response was incorrect: %+v", got)
	}
}

func TestStreamGetChannel_Online(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/streams/levi", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "_links": { "channel": "h", "self": "h" }, "stream": { "_links": { "self": "h" }, "_id": 1, "viewers": 1, "created_at": "2014-09-12T02:03:17Z", "preview": { "large": "l", "medium": "m", "small": "s", "template": "t" }, "channel": { "display_name": "t", "_links": { "stream_key": "h", "editors": "h", "subscriptions": "h", "commercial": "h", "videos": "h", "follows": "h", "self": "h", "chat": "h", "features": "h" }, "teams": [], "status": "t", "created_at": "2011-12-23T18:03:44Z", "logo": "h", "updated_at": "2013-02-15T15:22:24Z", "mature": null, "video_banner": null, "_id": 1, "background": "h", "banner": "h", "name": "t", "delay": 0, "url": "h", "game": "m" }, "game": "m" } }`)
	})
	want := &Stream{
		ID:        intPtr(1),
		CreatedAt: stringPtr("2014-09-12T02:03:17Z"),
		Viewers:   intPtr(1),
		Game:      stringPtr("m"),
		Channel: &Channel{
			ID:          intPtr(1),
			DisplayName: stringPtr("t"),
			Name:        stringPtr("t"),
			Game:        stringPtr("m"),
			Delay:       intPtr(0),
			Status:      stringPtr("t"),
			Teams:       make([]Team, 0),
			Banner:      stringPtr("h"),
			Background:  stringPtr("h"),
			Logo:        stringPtr("h"),
			URL:         stringPtr("h"),
			CreatedAt:   stringPtr("2011-12-23T18:03:44Z"),
			UpdatedAt:   stringPtr("2013-02-15T15:22:24Z"),
		},
		Preview: assetPtr(),
	}
	got, _, err := client.Streams.GetChannel("levi")
	if err != nil {
		t.Errorf("Streams.GetChannel returned error: %+v", err)
	}
	if !reflect.DeepEqual(got, want) {
		fmt.Printf("Streams.GetChannel Channel: %+v\n", *got.Channel)
		t.Errorf("Streams.GetChannel response was incorrect: \n want: %+v\n got:  %+v", want, got)
	}
}

func TestStreamGetChannel_Offline(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/streams/levi", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "_links": { "channel": "h", "self": "h" }, "stream": null } }`)
	})
	want := new(Stream)
	got, _, err := client.Streams.GetChannel("levi")
	if err != nil {
		t.Errorf("Streams.GetChannel returned error: %+v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Streams.GetChannel response was incorrect: %+v\n %+v", got, want)
	}
}

func TestStreamListStreams(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/streams", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"game":       "g",
			"channel":    "c",
			"limit":      "1",
			"offset":     "1",
			"embeddable": "true",
			"hls":        "true",
		})
		fmt.Fprint(w, `{ "_total": 1, "streams": [ { "_id": 1, "preview": { "medium": "m", "small": "s", "large": "l", "template": "t" }, "game": "g", "channel": { "mature": null, "background": "b", "updated_at": "2013-03-04T05:27:27Z", "_id": 1, "status": "s", "logo": "l", "teams": [], "url": "u", "display_name": "dn", "game": "g", "banner": "b", "name": "n", "delay": 0, "video_banner": null, "_links": { "chat": "h", "subscriptions": "h", "features": "h", "commercial": "h", "stream_key": "h", "editors": "h", "videos": "h", "self": "h", "follows": "h" }, "created_at": "2012-07-01T21:09:58Z" }, "viewers": 1, "created_at": "2014-09-12T02:03:17Z", "_links": { "self": "h" } } ], "_links": { "summary": "h", "followed": "h", "next": "https://api.twitch.tv/kraken/streams?channel=zisss%2Cvoyboy&game=Diablo+III&limit=100&offset=100", "featured": "h", "self": "https://api.twitch.tv/kraken/streams?channel=zisss%2Cvoyboy&game=Diablo+III&limit=100&offset=0" } }`)
	})

	want := []Stream{
		Stream{
			ID:        intPtr(1),
			Viewers:   intPtr(1),
			CreatedAt: stringPtr("2014-09-12T02:03:17Z"),
			Preview:   assetPtr(),
			Channel:   &Channel{},
			Game:      stringPtr("g"),
		},
	}

	opts := &StreamOptions{Game: "g", Channel: "c", Embeddable: true, RequestOptions: RequestOptions{Limit: 1, Offset: 1, HLS: true}}
	got, resp, err := client.Streams.ListStreams(opts)
	if err != nil {
		t.Errorf("Streams.ListStreams returned error: %v", err)
	}

	fmt.Printf("Streams.ListStreams: %+v\n", resp)

	if resp.Total != intPtr(1) {
		t.Errorf("Streams.ListStreams Total did not return correct value: %+v\n", resp.Total)
	}

	if resp.NextOffset != intPtr(100) {
		t.Errorf("Streams.ListStreams NextOffset did not return correct value: %+v\n", resp.NextOffset)
	}

	if resp.PrevOffset != nil {
		t.Errorf("Streams.ListStreams PrevOffset was not nil: %+v\n", *resp.PrevOffset)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Streams.ListStreams response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
