package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/users/test_user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "type": "t", "name": "n", "created_at": "2011-03-19T15:42:22Z", "updated_at": "2012-06-14T00:14:27Z", "logo": "l", "_id": 1, "display_name": "d", "bio": "b" }`)
	})
	want := &User{
		ID:          intPtr(1),
		Type:        stringPtr("t"),
		Name:        stringPtr("n"),
		Logo:        stringPtr("l"),
		DisplayName: stringPtr("d"),
		Bio:         stringPtr("b"),
		CreatedAt:   stringPtr("2011-03-19T15:42:22Z"),
		UpdatedAt:   stringPtr("2012-06-14T00:14:27Z"),
	}
	got, _, err := client.Users.GetUser("test_user")
	if err != nil {
		t.Errorf("Users.GetUser: request returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Users.GetUser response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestGetCurrentUser(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "type": "t", "name": "n", "created_at": "2011-06-03T17:49:19Z", "updated_at": "2012-06-18T17:19:57Z", "logo": "l", "_id": 1, "display_name": "d", "email": "e", "partnered": true, "bio": "b" }`)
	})
	want := &User{
		ID:          intPtr(1),
		Type:        stringPtr("t"),
		Name:        stringPtr("n"),
		Logo:        stringPtr("l"),
		DisplayName: stringPtr("d"),
		Bio:         stringPtr("b"),
		Email:       stringPtr("e"),
		Partnered:   boolPtr(true),
		CreatedAt:   stringPtr("2011-06-03T17:49:19Z"),
		UpdatedAt:   stringPtr("2012-06-18T17:19:57Z"),
	}
	got, _, err := client.Users.GetCurrentUser()
	if err != nil {
		t.Errorf("Users.GetCurrentUser: request returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Users.GetCurrentUser response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestListFollowedStreams(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/streams/followed", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":  "1",
			"offset": "1",
			"hls":    "true",
		})
		fmt.Fprint(w, `{ "_total": 1, "streams": [ { "_id": 1, "preview": { "medium": "m", "small": "s", "large": "l", "template": "t" }, "game": "g", "channel": { "mature": null, "background": "b", "updated_at": "2013-02-15T15:22:24Z", "_id": 1, "status": "s", "logo": "l", "teams": [], "url": "u", "display_name": "d", "game": "g", "banner": "b", "name": "n", "delay": 0, "video_banner": null, "_links": { "chat": "c", "subscriptions": "s", "features": "f", "commercial": "c", "stream_key": "s", "editors": "e", "videos": "v", "self": "s", "follows": "f" }, "created_at": "2011-12-23T18:03:44Z" }, "viewers": 1, "created_at": "2014-09-12T02:03:17Z", "_links": { "self": "h" } } ], "_links": { "summary": "h", "followed": "h", "next": "https://api.twitch.tv/kraken/streams?channel=zisss%2Cvoyboy&game=Diablo+III&limit=100&offset=100", "featured": "f", "self": "https://api.twitch.tv/kraken/streams?channel=zisss%2Cvoyboy&game=Diablo+III&limit=100&offset=0" } }`)
	})

	opts := &RequestOptions{HLS: true, ListOptions: ListOptions{Limit: 1, Offset: 1}}
	got, resp, err := client.Users.ListFollowedStreams(opts)
	if err != nil {
		t.Errorf("Users.ListFollowedStreams: returned error: %v", err)
	}

	testListResponse(t, resp, intPtr(1), intPtr(100), nil)
	want := []Stream{
		Stream{
			ID:        intPtr(1),
			Viewers:   intPtr(1),
			CreatedAt: stringPtr("2014-09-12T02:03:17Z"),
			Preview:   assetPtr(),
			Channel:   channelPtr(),
			Game:      stringPtr("g"),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Users.ListFollowedStreams response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestListFollowedVideos(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/videos/followed", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":  "1",
			"offset": "1",
		})
		fmt.Fprint(w, `{ "_links": { "next": "https://api.twitch.tv/kraken/videos/top?game=League+of+Legends&limit=10&offset=10&period=month", "self": "https://api.twitch.tv/kraken/videos/top?game=League+of+Legends&limit=10&offset=0&period=month" }, "videos": [ { "recorded_at": "2013-03-13T09:51:31Z", "preview": "p", "description": "d", "url": "u", "title": "t", "channel": { "name": "n", "display_name": "d" }, "length": 1, "game": "g", "views": 1, "_id": "i", "_links": { "channel": "c", "self": "s" } } ] }`)
	})

	opts := &ListOptions{Limit: 1, Offset: 1}
	got, resp, err := client.Users.ListFollowedVideos(opts)
	if err != nil {
		t.Errorf("Users.ListFollowedVideos: returned error: %v", err)
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
		t.Errorf("Users.ListFollowedVideos response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
