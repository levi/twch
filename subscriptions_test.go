package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetChannelSubscriptions(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/channels/test_channel/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testParams(t, r, params{
			"limit":     "1",
			"offset":    "1",
			"direction": "desc",
		})
		fmt.Fprint(w, `{  "_total": 1,  "_links": {    "next": "https://api.twitch.tv/kraken/channels/test_channel/subscriptions?limit=25&offset=25",    "self": "https://api.twitch.tv/kraken/channels/test_channel/subscriptions?limit=25&offset=0"  },  "subscriptions": [    {      "_id": "i",      "user": {        "_id": 1,        "logo": null,        "staff": false,        "created_at": "2012-12-06T00:32:36Z",        "name": "n",        "updated_at": "2013-02-06T21:27:46Z",        "display_name": "d",        "_links": {          "self": "s"        }      },      "created_at": "2013-02-06T21:33:33Z",      "_links": {        "self": "s"      }    }  ]}`)
	})
	want := []Subscription{
		Subscription{
			ID: stringPtr("i"),
			User: &User{
				ID:          intPtr(1),
				Staff:       boolPtr(false),
				CreatedAt:   stringPtr("2012-12-06T00:32:36Z"),
				Name:        stringPtr("n"),
				UpdatedAt:   stringPtr("2013-02-06T21:27:46Z"),
				DisplayName: stringPtr("d"),
			},
			CreatedAt: stringPtr("2013-02-06T21:33:33Z"),
		},
	}
	opts := &SubscriptionOptions{Direction: "desc", ListOptions: ListOptions{Limit: 1, Offset: 1}}
	got, resp, err := client.Subscriptions.GetChannelSubscriptions("test_channel", opts)
	if err != nil {
		t.Errorf("Subscriptions.GetChannelSubscriptions: request returned error %+v", err)
	}

	testListResponse(t, resp, intPtr(1), intPtr(25), nil)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Subscriptions.GetChannelSubscriptions response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func TestGetUserSubscribed(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/channels/test_channel/subscriptions/test_user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "_id": "i", "user": { "_id": 1, "logo": null, "staff": false, "created_at": "2012-12-06T00:32:36Z", "name": "n", "updated_at": "2013-02-06T21:27:46Z", "display_name": "d", "_links": { "self": "s" } }, "created_at": "2013-02-06T21:33:33Z", "_links": { "self": "s" } }`)
	})
	want := &Subscription{
		ID: stringPtr("i"),
		User: &User{
			ID:          intPtr(1),
			Staff:       boolPtr(false),
			CreatedAt:   stringPtr("2012-12-06T00:32:36Z"),
			Name:        stringPtr("n"),
			UpdatedAt:   stringPtr("2013-02-06T21:27:46Z"),
			DisplayName: stringPtr("d"),
		},
		CreatedAt: stringPtr("2013-02-06T21:33:33Z"),
	}
	got, resp, err := client.Subscriptions.GetUserSubscribed("test_channel", "test_user")
	if err != nil {
		t.Errorf("Subscriptions.GetUserSubscribed: request returned error %+v", err)
	}

	testListResponse(t, resp, nil, nil, nil)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Subscriptions.GetUserSubscribed response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
