package twch

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListEmoticons(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/chat/emoticons", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{  "_links": {    "self": "s"  },  "emoticons": [    {      "regex": "r",      "images": [        {          "emoticon_set": 1,          "height": 1,          "width": 1,          "url": "u"        }      ]    }  ]}`)
	})
	want := []Emoticon{
		Emoticon{
			Regex: stringPtr("r"),
			Images: []EmoticonImage{
				EmoticonImage{
					EmoticonSet: intPtr(1),
					Height:      intPtr(1),
					Width:       intPtr(1),
					URL:         stringPtr("u"),
				},
			},
		},
	}
	got, _, err := client.Chat.ListEmoticons()
	if err != nil {
		t.Errorf("Chat.ListEmoticons: request returned error %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Chat.ListEmoticons response did not match:\nwant: %+v\ngot:  %+v", want, got)
	}
}
