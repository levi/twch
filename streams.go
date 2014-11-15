package twch

import (
  "fmt"
)

type Streams struct {
  client *Client
}

type streamSummary struct {
  Viewers  int         `json:"viewers"`
  Channels int         `json:"channels"`
  Links    interface{} `json:"_links"`
}

type streamChannel struct {
  Links  interface{} `json:"_links"`
  Stream Stream      `json:"stream"`
}

type streamList struct {
  Total   int         `json:"_total"`
  Streams []Stream    `json:"stream"`
  Links   interface{} `json:"_links"`
}

type streamFeatured struct {
  Featured []Stream    `json:"featured"`
  Links    interface{} `json:"_links"`
}

type streamFollowed struct {
  Links   interface{} `json:"_links"`
  Total   int         `json:"_total"`
  Streams []Stream    `json:"streams"`
}

type Stream struct {
  Id        int         `json:"_id"`
  CreatedAt string      `json:"created_at"`
  Preview   Asset       `json:"preview"`
  Channel   interface{} `json:"channel"`
  Game      string      `json:"game"`
}

type RequestOptions struct {
  limit  int
  offset int
  hls    bool
}

type StreamsRequestOptions struct {
  RequestOptions
  game       string
  channel    string
  embeddable bool
  client_id  string
}

func (s *Streams) Channel(channel string) (stream Stream, err error) {
  uri := fmt.Sprintf("streams/%s", channel)
  req, err := s.client.NewRequest("GET", uri)
  if err != nil {
    return nil, err
  }

  r := new()
}

func (s *Streams) Summary() (viewers, channels int, err error) {
  req, err := s.client.NewRequest("GET", "streams/summary")
  if err != nil {
    return 0, 0, err
  }

  r := new(streamSummary)
  _, err = s.client.Do(req, r)
  if err != nil {
    return 0, 0, err
  }

  return r.Viewers, r.Channels, nil
}
