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

func TestStreamGetStream_Online(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/streams/levi", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "_links": { "channel": "h", "self": "h" }, "stream": { "_links": { "self": "h" }, "_id": 1, "viewers": 1, "created_at": "2014-09-12T02:03:17Z", "preview": { "large": "l", "medium": "m", "small": "s", "template": "t" }, "channel": { "display_name": "d", "_links": { "stream_key": "h", "editors": "h", "subscriptions": "h", "commercial": "h", "videos": "h", "follows": "h", "self": "h", "chat": "h", "features": "h" }, "teams": [], "status": "s", "created_at": "2011-12-23T18:03:44Z", "logo": "l", "updated_at": "2013-02-15T15:22:24Z", "mature": null, "video_banner": null, "_id": 1, "background": "b", "banner": "b", "name": "n", "delay": 0, "url": "u", "game": "g" }, "game": "g" } }`)
	})
	want := &Stream{
		ID:        intPtr(1),
		CreatedAt: stringPtr("2014-09-12T02:03:17Z"),
		Viewers:   intPtr(1),
		Game:      stringPtr("g"),
		Channel:   channelPtr(),
		Preview:   assetPtr(),
	}
	got, _, err := client.Streams.GetStream("levi")
	if err != nil {
		t.Errorf("Streams.GetStream returned error: %+v", err)
	}
	if !reflect.DeepEqual(got, want) {
		fmt.Printf("Streams.GetStream Channel: %+v\n", *got.Channel)
		t.Errorf("Streams.GetStream response was incorrect: \n want: %+v\n got:  %+v", want, got)
	}
}

func TestStreamGetGetStream_Offline(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/streams/levi", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "_links": { "channel": "h", "self": "h" }, "stream": null } }`)
	})
	want := new(Stream)
	got, _, err := client.Streams.GetStream("levi")
	if err != nil {
		t.Errorf("Streams.GetStream returned error: %+v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Streams.GetStream response was incorrect: %+v\n %+v", got, want)
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
			"client_id":  "c",
		})
		fmt.Fprint(w, `{"streams":[{"_id":1,"game":"g","viewers":1,"created_at":"2014-11-21T07:09:54Z","preview":{"small":"s","medium":"m","large":"l","template":"t"},"channel":{"background":null,"banner":null,"display_name":"d","game":"g","logo":"l","mature":false,"status":"s","partner":false,"url":"u","video_banner":null,"_id":1,"name":"n","created_at":"2011-12-23T18:03:44Z","updated_at":"2013-02-15T15:22:24Z","delay":0,"followers":1,"profile_banner":null,"profile_banner_background_color":null,"views":1,"language":"en"}}],"_total":1,"_links":{"self":"https://api.twitch.tv/kraken/streams?game=Diablo+III&limit=25&offset=0","next":"https://api.twitch.tv/kraken/streams?game=Diablo+III&limit=25&offset=25","featured":"https://api.twitch.tv/kraken/streams/featured","summary":"https://api.twitch.tv/kraken/streams/summary","followed":"https://api.twitch.tv/kraken/streams/followed"}}`)
	})

	want := []Stream{
		Stream{
			ID:        intPtr(1),
			Viewers:   intPtr(1),
			CreatedAt: stringPtr("2014-11-21T07:09:54Z"),
			Preview:   assetPtr(),
			Channel: &Channel{
				ID:          intPtr(1),
				DisplayName: stringPtr("d"),
				Name:        stringPtr("n"),
				Game:        stringPtr("g"),
				Delay:       intPtr(0),
				Status:      stringPtr("s"),
				Logo:        stringPtr("l"),
				URL:         stringPtr("u"),
				Mature:      boolPtr(false),
				Partner:     boolPtr(false),
				Followers:   intPtr(1),
				Views:       intPtr(1),
				Language:    stringPtr("en"),
				CreatedAt:   stringPtr("2011-12-23T18:03:44Z"),
				UpdatedAt:   stringPtr("2013-02-15T15:22:24Z"),
			},
			Game: stringPtr("g"),
		},
	}

	opts := &StreamOptions{Game: "g", Channel: "c", Embeddable: true, ClientID: "c", RequestOptions: RequestOptions{ListOptions: ListOptions{Limit: 1, Offset: 1}, HLS: true}}
	got, resp, err := client.Streams.ListStreams(opts)
	if err != nil {
		t.Errorf("Streams.ListStreams returned error: %v", err)
	}
	testListResponse(t, resp, intPtr(1), intPtr(25), nil)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Streams.ListStreams response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestStreamListFeaturedStreams(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/streams/featured", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":  "1",
			"offset": "1",
			"hls":    "true",
		})
		fmt.Fprint(w, `{  "_links": {     "self": "https://api.twitch.tv/kraken/streams/featured?limit=25&offset=0",     "next": "https://api.twitch.tv/kraken/streams/featured?limit=25&offset=25"  },  "featured": [    {      "image": "i",      "text": "t",      "stream": {        "_id":1,        "game":"g",        "viewers":1,        "created_at":"2014-11-18T16:36:57Z",        "preview":{          "small":"s","medium":"m","large":"l","template":"t"        },        "channel":{          "mature":true,          "status":"s",          "broadcaster_language":"en",          "display_name":"d",          "game":"g",          "delay":0,          "language":"en-us",          "_id":1,          "name":"n",          "created_at":"2008-04-27T11:40:03Z",          "updated_at":"2014-11-18T17:02:30Z",          "logo":"l","banner":"b",          "video_banner":"b",          "background":null,"profile_banner":"p",          "profile_banner_background_color":"p",          "partner":true,          "url":"u",          "views":1,          "followers":1        }      }    }  ]}`)
	})

	want := []FeaturedStream{
		FeaturedStream{
			Image: stringPtr("i"),
			Text:  stringPtr("t"),
			Stream: &Stream{
				ID:        intPtr(1),
				Game:      stringPtr("g"),
				Viewers:   intPtr(1),
				CreatedAt: stringPtr("2014-11-18T16:36:57Z"),
				Preview:   assetPtr(),
				Channel: &Channel{
					ID:                           intPtr(1),
					DisplayName:                  stringPtr("d"),
					Name:                         stringPtr("n"),
					Game:                         stringPtr("g"),
					BroadcasterLanguage:          stringPtr("en"),
					Language:                     stringPtr("en-us"),
					Delay:                        intPtr(0),
					Status:                       stringPtr("s"),
					Banner:                       stringPtr("b"),
					VideoBanner:                  stringPtr("b"),
					ProfileBanner:                stringPtr("p"),
					ProfileBannerBackgroundColor: stringPtr("p"),
					Mature:    boolPtr(true),
					Partner:   boolPtr(true),
					Views:     intPtr(1),
					Followers: intPtr(1),
					Logo:      stringPtr("l"),
					URL:       stringPtr("u"),
					CreatedAt: stringPtr("2008-04-27T11:40:03Z"),
					UpdatedAt: stringPtr("2014-11-18T17:02:30Z"),
				},
			},
		},
	}

	opts := &RequestOptions{ListOptions: ListOptions{Limit: 1, Offset: 1}, HLS: true}
	got, resp, err := client.Streams.ListFeaturedStreams(opts)
	if err != nil {
		t.Errorf("Streams.ListFeaturedStreams returned error: %v", err)
	}
	testListResponse(t, resp, nil, intPtr(25), nil)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Streams.ListFeaturedStreams response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
