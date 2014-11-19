package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListIngests(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/ingests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{  "_links": {    "self": "h"  },  "ingests": [    {      "name": "n" ,      "default": false,      "_id": 1,      "url_template": "u",      "availability":1.0    }  ]}`)
	})

	got, resp, err := client.Ingests.ListIngests()
	if err != nil {
		t.Errorf("Ingests.ListIngests: returned error: %v", err)
	}

	testListResponse(t, resp, nil, nil, nil)
	want := []Ingest{
		Ingest{
			Name:         stringPtr("n"),
			Default:      boolPtr(false),
			ID:           intPtr(1),
			URLTemplate:  stringPtr("u"),
			Availability: float32Ptr(1.0),
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Ingests.ListIngests: response was wrong: %+v", got)
	}
}
