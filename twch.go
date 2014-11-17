package twch

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
	"strconv"
)

const (
	baseUrl      = "https://api.twitch.tv/kraken/"
	acceptHeader = "application/vnd.twitchtv.v3+json"
)

type Client struct {
	client *http.Client

	ID string

	BaseUrl *url.URL

	Blocks        *Blocks
	Channels      *Channels
	Chat          *Chats
	Follows       *Follows
	Games         *Games
	Ingests       *Ingests
	Search        *Search
	Streams       *Streams
	Subscriptions *Subscriptions
	Teams         *Teams
	Users         *Users
	Videos        *Videos
}

// NewClient constructs a new client to interface with the Twitch API
func NewClient(id string) (client *Client, err error) {
	client = new(Client)
	client.ID = id
	client.client = http.DefaultClient
	client.BaseUrl, _ = url.Parse(baseUrl)
	client.Blocks = &Blocks{client: client}
	client.Channels = &Channels{client: client}
	client.Chat = &Chats{client: client}
	client.Follows = &Follows{client: client}
	client.Games = &Games{client: client}
	client.Ingests = &Ingests{client: client}
	client.Search = &Search{client: client}
	client.Streams = &Streams{client: client}
	client.Subscriptions = &Subscriptions{client: client}
	client.Teams = &Teams{client: client}
	client.Users = &Users{client: client}
	client.Videos = &Videos{client: client}
	return client, nil
}

// appendOptions creates a relative URL string that includes query params provided as a struct
func appendOptions(u string, opts interface{}) (s string, err error) {
	url, err := url.Parse(u)
	if err != nil {
		return
	}

	v, err := query.Values(opts)
	if err != nil {
		return
	}

	url.RawQuery = v.Encode()

	return url.String(), nil
}

// NewRequest constructs a valid http.Request object for Twitch requests
func (c *Client) NewRequest(method, uri string) (req *http.Request, err error) {
	apiUri, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	reqUrl := c.BaseUrl.ResolveReference(apiUri)
	q, err := url.ParseQuery(reqUrl.RawQuery)
	if err != nil {
		return nil, err
	}
	q.Add("client_id", c.ID)
	reqUrl.RawQuery = q.Encode()

	req, err = http.NewRequest(method, reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", acceptHeader)

	return req, nil
}

// Do performs the http request and marshals the response JSON into the past `v` interface type.
func (c *Client) Do(req *http.Request, v interface{}) (r *Response, err error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return nil, err
	}

	// TODO: Include next/prev and total meta data
	r, err = newResponse(resp, nil, nil)
	if err != nil {
		return
	}

	return r, nil
}

type listTotalOptions interface {
	ListTotal() *int
}

type listPageOptions interface {
	NextOffset() (*int, error)
	PrevOffset() (*int, error)
}

type listTotal struct {
	Total *int `json:"_total"`
}

func (l *listTotal) ListTotal() *int {
	return l.Total
}

// listLinks is an abstract representation of response paging links
type listLinks struct {
	Links listPagingLinks `json:"_links"`
}

// listPagingLinks i
type listPagingLinks struct {
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

// NextOffset returns the offset count needed for the next list query.
// A nil pointer is returned if the response never included data for the next offset.
func (l *listLinks) NextOffset() (*int, error) {
	return urlOffsetVal(l.Links.Next)
}

// PrevOffset returns the offset count needed for the prev list query
// A nil pointer is returned if the response never included data for the next offset.
func (l *listLinks) PrevOffset() (*int, error) {
	return urlOffsetVal(l.Links.Prev)
}

// urlOffsetVal extracts the "offset" query string value from an URL string representation.
// The returned value is converted into an integer. If the URL fails to contain the
// "offset" parameter, a nil pointer is returned.
func urlOffsetVal(s string) (*int, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	qs, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}

	if o := qs.Get("offset"); o != "" {
		i, err := strconv.Atoi(o)
		if err != nil {
			return nil, err
		}

		v := new(int)
		*v = i
		return v, nil
	}

	return nil, nil
}

// Asset represents links to images assets that are likely to come along with Game and Channel responses.
// Fields are pointer types to support empty responses from API results.
type Asset struct {
	Large    *string `json:"large"`
	Medium   *string `json:"medium"`
	Small    *string `json:"small"`
	Template *string `json:"template"`
}

// RequestOptions is the base query parameters used for customizing query output from List queries.
type RequestOptions struct {
	Limit  int  `url:"limit,omitempty"`
	Offset int  `url:"offset,omitempty"`
	HLS    bool `url:"hls,omitempty"`
}

// Response augments http.Response to include extra meta data for List query responses.
// Fields are pointer types to support response inconsistencies in Twitch's API, since some endpoints
// allow pagination, but don't include a total count, etc.
type Response struct {
	NextOffset *int
	PrevOffset *int
	Total      *int
	*http.Response
}

// newResponse constructs a new response with included paging meta data.
// Metadata is completely optional in the construction of the Response and passing
// nil pointers will keep the final fields nil.
func newResponse(resp *http.Response, p listPageOptions, t listTotalOptions) (r *Response, err error) {
	r = &Response{Response: resp}

	if p != nil {
		r.NextOffset, err = p.NextOffset()
		if err != nil {
			return
		}

		r.PrevOffset, err = p.PrevOffset()
		if err != nil {
			return
		}
	}

	if t != nil {
		r.Total = t.ListTotal()
	}

	return r, nil
}

// intPtr converts an int value into an allocated pointer to an int
func intPtr(n int) *int {
	i := new(int)
	*i = n
	return i
}

// stringPtr converts a string value into an allocated pointer to a string
func stringPtr(str string) *string {
	s := new(string)
	*s = str
	return s
}
