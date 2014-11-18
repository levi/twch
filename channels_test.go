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

func TestGetUserChannel(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/channel", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{  "game": "g",  "name": "n",  "stream_key": "s",  "created_at": "2011-03-19T15:42:22Z",  "delay": 0,  "status": "s",  "updated_at": "2012-03-14T03:30:41Z",  "teams": [{    "name": "n",    "created_at": "2011-10-25T23:55:47Z",    "updated_at": "2011-11-14T19:48:21Z",    "background": null,    "banner": "b",    "logo": null,    "_links": {      "self": "s"    },    "_id": 10,    "info": "i",    "display_name": "d"  }],  "_links": {    "self": "s",    "chat":"c",    "videos": "v",    "video_status": "v",    "commercial":"c"  },  "banner": "b",  "video_banner": "v",  "background": "b",  "logo": "l",  "_id": 1,  "mature": false,  "login": "l",  "url": "u",  "email": "e"}`)
	})
	want := &Channel{
		// Title:       stringPtr("t"),
		// DisplayName: stringPtr("d"),
		Name:        stringPtr("n"),
		Game:        stringPtr("g"),
		StreamKey:   stringPtr("s"),
		Status:      stringPtr("s"),
		Delay:       intPtr(0),
		Banner:      stringPtr("b"),
		VideoBanner: stringPtr("v"),
		Background:  stringPtr("b"),
		Logo:        stringPtr("l"),
		ID:          intPtr(1),
		Mature:      boolPtr(false),
		URL:         stringPtr("u"),
		Email:       stringPtr("e"),
		Login:       stringPtr("l"),
		Teams: []Team{
			Team{
				Name:        stringPtr("n"),
				Banner:      stringPtr("b"),
				ID:          intPtr(10),
				Info:        stringPtr("i"),
				DisplayName: stringPtr("d"),
				CreatedAt:   stringPtr("2011-10-25T23:55:47Z"),
				UpdatedAt:   stringPtr("2011-11-14T19:48:21Z"),
			},
		},
		CreatedAt: stringPtr("2011-03-19T15:42:22Z"),
		UpdatedAt: stringPtr("2012-03-14T03:30:41Z"),
	}
	got, _, err := client.Channels.GetUserChannel()
	if err != nil {
		t.Errorf("Channels.GetUserChannel: request returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Channels.GetUserChannel response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestListChannelEditors(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/channels/test_user1/editors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{  "_links": {    "self": "s"  },  "users": [    {      "_links": {        "self": "s"      },      "created_at": "2013-02-06T21:21:57Z",      "name": "n",      "updated_at": "2013-02-13T20:59:42Z",      "_id": 1,      "display_name": "d",      "logo": null,      "staff": false    }  ]}`)
	})
	want := []User{
		User{
			ID:          intPtr(1),
			DisplayName: stringPtr("d"),
			Name:        stringPtr("n"),
			Staff:       boolPtr(false),
			CreatedAt:   stringPtr("2013-02-06T21:21:57Z"),
			UpdatedAt:   stringPtr("2013-02-13T20:59:42Z"),
		},
	}
	got, _, err := client.Channels.ListChannelEditors("test_user1")
	if err != nil {
		t.Errorf("Channels.ListChannelEditors: request returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Channels.ListChannelEditors response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
