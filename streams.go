package twch

import (
	"fmt"
)

type Streams struct {
	client *Client
}

type streamChannel struct {
	Links  interface{} `json:"_links,omitempty"`
	Stream Stream      `json:"stream,omitempty"`
}

type streamList struct {
	Streams []Stream    `json:"streams,omitempty"`
	*listLinks
	*listTotal
}

type streamFeatured struct {
	Featured []Stream    `json:"featured,omitempty"`
	Links    interface{} `json:"_links,omitempty"`
}

type streamFollowed struct {
	Links   interface{} `json:"_links,omitempty"`
	Total   int         `json:"_total,omitempty"`
	Streams []Stream    `json:"streams,omitempty"`
}

type StreamSummary struct {
	Viewers  int `json:"viewers,omitempty"`
	Channels int `json:"channels,omitempty"`
}

type Stream struct {
	ID        *int     `json:"_id,omitempty"`
	Viewers   *int     `json:"viewers,omitempty"`
	CreatedAt *string  `json:"created_at,omitempty"`
	Preview   *Asset   `json:"preview,omitempty"`
	Channel   *Channel `json:"channel,omitempty"`
	Game      *string  `json:"game,omitempty"`
}

type StreamOptions struct {
	Game       string `url:"game,omitempty"`
	Channel    string `url:"channel,omitempty"`
	Embeddable bool   `url:"embeddable,omitempty"`
	RequestOptions
}

// Summary returns viewership and channel count for all streams currently on Twitch
func (s *Streams) GetSummary() (summary *StreamSummary, resp *Response, err error) {
	req, err := s.client.NewRequest("GET", "streams/summary")
	if err != nil {
		return
	}

	r := new(StreamSummary)
	resp, err = s.client.Do(req, r)
	if err != nil {
		return
	}

	return r, resp, nil
}

func (s *Streams) ListStreams(opts *StreamOptions) (streams []Stream, resp *Response, err error) {
	u, err := appendOptions("streams", opts)
	if err != nil {
		return
	}

	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return
	}

	r := new(streamList)
	resp, err = s.client.Do(req, r)
	if err != nil {
		return
	}
	streams = r.Streams
	return
}

// GetChannel fetches the current stream of a given channel.
// If the channel is offline, a zeroed Stream is returned without error.
func (s *Streams) GetChannel(channel string) (stream *Stream, resp *Response, err error) {
	uri := fmt.Sprintf("streams/%s", channel)
	req, err := s.client.NewRequest("GET", uri)
	if err != nil {
		return
	}

	r := new(streamChannel)
	resp, err = s.client.Do(req, r)
	if err != nil {
		return
	}

	return &r.Stream, resp, nil
}
