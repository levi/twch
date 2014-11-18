package twch

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client, _ = NewClient("test-client-key")
	client.BaseUrl, _ = url.Parse(server.URL)
}

func teardown() {
	server.Close()
}

type params map[string]string

// testParams compares the request's query params and body query with a custom defined map of string key values
func testParams(t *testing.T, r *http.Request, par params) {
	want := url.Values{}
	for k, v := range par {
		want.Add(k, v)
	}
	want.Add("client_id", client.ID)
	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("testParams: Request params failed to match:\nwant: %+v\ngot:  %+v", want, got)
	}
}

func testMethod(t *testing.T, r *http.Request, method string) {
	if r.Method != method {
		t.Errorf("Top Games: method = %s, want %v", r.Method, method)
	}
}

func testResponse(t *testing.T, r *Response) {
}

func testListResponse(t *testing.T, r *Response, total *int, next *int, prev *int) {
	if total != nil {
		if r.Total == nil {
			t.Errorf("ListResponse: response total is nil")
		} else if *r.Total != *total {
			t.Errorf("ListResponse: response total should not include value: %+v", *r.Total)
		}
	} else {
		if r.Total != total {
			t.Errorf("ListResponse: response total should not include value: %+v", *r.Total)
		}
	}

	if next != nil {
		if r.NextOffset == nil {
			t.Errorf("ListResponse: response next is nil")
		} else if *r.NextOffset != *next {
			t.Errorf("ListResponse: response next should not include value: %+v", *r.NextOffset)
		}
	} else {
		if r.NextOffset != next {
			t.Errorf("ListResponse: response next should not include value: %+v", *r.NextOffset)
		}
	}

	if prev != nil {
		if r.PrevOffset == nil {
			t.Errorf("ListResponse: response prev is nil")
		} else if *r.PrevOffset != *prev {
			t.Errorf("ListResponse: response prev should not include value: %+v", *r.PrevOffset)
		}
	} else {
		if r.PrevOffset != prev {
			t.Errorf("ListResponse: response prev should not include value: %+v", *r.PrevOffset)
		}
	}
}

func assetPtr() *Asset {
	return &Asset{
		Large:    stringPtr("l"),
		Medium:   stringPtr("m"),
		Small:    stringPtr("s"),
		Template: stringPtr("t"),
	}
}

func channelPtr() *Channel {
	return &Channel{
		ID:          intPtr(1),
		DisplayName: stringPtr("d"),
		Name:        stringPtr("n"),
		Game:        stringPtr("g"),
		Delay:       intPtr(0),
		Status:      stringPtr("s"),
		Teams:       make([]Team, 0),
		Banner:      stringPtr("b"),
		Background:  stringPtr("b"),
		Logo:        stringPtr("l"),
		URL:         stringPtr("u"),
		CreatedAt:   stringPtr("2011-12-23T18:03:44Z"),
		UpdatedAt:   stringPtr("2013-02-15T15:22:24Z"),
	}
}

func TestListTotal(t *testing.T) {
	l := listTotal{Total: intPtr(1)}
	if v := l.ListTotal(); *v != 1 {
		t.Errorf("listTotal.Total() did not return correct number. Got: %v", *v)
	}
}

func TestListLinks(t *testing.T) {
	l := listLinks{
		Links: listPagingLinks{
			Next: "https://api.twitch.tv/kraken/streams?channel=zisss%2Cvoyboy&game=Diablo+III&limit=100&offset=300",
			Prev: "https://api.twitch.tv/kraken/streams?channel=zisss%2Cvoyboy&game=Diablo+III&limit=100&offset=100",
		},
	}
	n, err := l.NextOffset()
	if err != nil {
		t.Errorf("ListLinks.NextOffset() returned an error: %+v\n", err)
	}

	if *n != 300 {
		t.Errorf("listLinks.NextOffset() did not return correct value. Got: %v\n", *n)
	}

	p, err := l.PrevOffset()
	if err != nil {
		t.Errorf("ListLinks.PrevOffset() returned an error: %+v\n", err)
	}

	if *p != 100 {
		t.Errorf("listLinks.NextOffset() did not return correct value. Got: %v\n", *p)
	}
}

func TestListLinks_empty(t *testing.T) {
	l := listLinks{}
	n, err := l.NextOffset()
	if err != nil {
		t.Errorf("ListLinks.NextOffset() returned an error: %+v\n", err)
	}

	if n != nil {
		t.Errorf("listLinks.NextOffset() did not return correct value. Got: %v\n", *n)
	}

	p, err := l.PrevOffset()
	if err != nil {
		t.Errorf("ListLinks.PrevOffset() returned an error: %+v\n", err)
	}

	if p != nil {
		t.Errorf("listLinks.NextOffset() did not return correct value. Got: %v\n", *p)
	}
}

type testResponseData struct {
	*listLinks
	*listTotal
}

func TestNewResponse(t *testing.T) {
	resp := new(http.Response)
	want := &Response{
		Response:   resp,
		NextOffset: intPtr(300),
		PrevOffset: intPtr(100),
		Total:      intPtr(1),
	}

	got, err := newResponse(resp, testResponseData{
		listLinks: &listLinks{
			Links: listPagingLinks{
				Next: "https://api.twitch.tv/kraken/streams?limit=100&offset=300",
				Prev: "https://api.twitch.tv/kraken/streams?limit=100&offset=100",
			},
		},
		listTotal: &listTotal{Total: intPtr(1)},
	})

	if err != nil {
		t.Errorf("newResponse returned with an error: %+v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("newResponse did not return correct value\nwant: %+v\n got: %+v", want, got)
	}
}
