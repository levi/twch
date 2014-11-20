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
