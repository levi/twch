package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListBlocks(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/test_user1/blocks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{ "_links": { "next": "https://api.twitch.tv/kraken/users/test_user1/test_user1?limit=25&offset=25", "self": "https://api.twitch.tv/kraken/users/test_user1/test_user1?limit=25&offset=0" }, "blocks": [ { "_links": { "self": "s" }, "updated_at": "2013-02-07T01:04:43Z", "user": { "_links": { "self": "s" }, "updated_at": "2013-02-06T22:44:19Z", "display_name": "d", "staff": false, "name": "n", "_id": 1, "logo": "l", "created_at": "2010-06-30T08:26:49Z" }, "_id": 1 } ] }`)
	})

	want := []Block{
		Block{
			ID:        1,
			UpdatedAt: "2013-02-07T01:04:43Z",
			User: User{
				ID:          1,
				DisplayName: "d",
				Name:        "n",
				Logo:        "l",
				Staff:       boolPtr(false),
				CreatedAt:   "2010-06-30T08:26:49Z",
				UpdatedAt:   "2013-02-06T22:44:19Z",
			},
		},
	}

	opts := &ListOptions{Limit: 25, Offset: 0}
	got, resp, err := client.Blocks.ListBlocks("test_user1", opts)
	if err != nil {
		t.Errorf("Blocks.ListBlocks: returned an error: %+v\n", err)
	}

	testListResponse(t, resp, nil, intPtr(25), nil)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Block.ListBlocks: result did not match expecation\nwant: %+v\n got: %+v", want, got)
	}
}

func TestAddBlock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/test_user1/blocks/test_user2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		fmt.Fprint(w, `{ "_links": { "self": "h" }, "updated_at": "2013-02-07T01:04:43Z", "user": { "_links": { "self": "h" }, "updated_at": "2013-01-18T22:33:55Z", "logo": "l", "staff": false, "display_name": "d", "name": "n", "_id": 1, "created_at": "2011-05-01T14:50:12Z" }, "_id": 1 }`)
	})

	want := &Block{
		ID:        1,
		UpdatedAt: "2013-02-07T01:04:43Z",
		User: User{
			ID:          1,
			DisplayName: "d",
			Name:        "n",
			Logo:        "l",
			Staff:       boolPtr(false),
			CreatedAt:   "2011-05-01T14:50:12Z",
			UpdatedAt:   "2013-01-18T22:33:55Z",
		},
	}

	got, _, err := client.Blocks.AddBlock("test_user1", "test_user2")
	if err != nil {
		t.Errorf("Blocks.AddBlock: returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Block.AddBlocks: result did not match expecation\nwant: %+v\n got: %+v", want, got)
	}
}

func TestRemoveBlock(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/test_user1/blocks/test_user2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	err := client.Blocks.RemoveBlock("test_user1", "test_user2")
	if err != nil {
		t.Errorf("Blocks.RemoveBlock: returned error %+v", err)
	}
}
