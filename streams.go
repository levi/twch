package twch

type Streams struct {
  client *Client
}

type streamSummary struct {
  Viewers  int         `json:"viewers,omitempty"`
  Channels int         `json:"channels,omitempty"`
  Links    interface{} `json:"_links,omitempty"`
}

type streamChannel struct {
  Links  interface{} `json:"_links,omitempty"`
  Stream Stream      `json:"stream,omitempty"`
}

type streamList struct {
  Total   int         `json:"_total,omitempty"`
  Streams []Stream    `json:"stream,omitempty"`
  Links   interface{} `json:"_links,omitempty"`
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

type Stream struct {
  Id        int         `json:"_id,omitempty"`
  CreatedAt string      `json:"created_at,omitempty"`
  Preview   Asset       `json:"preview,omitempty"`
  Channel   interface{} `json:"channel,omitempty"`
  Game      string      `json:"game,omitempty"`
}

type RequestOptions struct {
  Limit  int  `url:"limit,omitempty"`
  Offset int  `url:"offset,omitempty"`
  HLS    bool `url:"hls,omitempty"`
}

type StreamOptions struct {
  RequestOptions
  Game       string `url:"game:omitempty"`
  Channel    string `url:"channel,omitempty"`
  Embeddable bool   `url:"embeddable,omitempty"`
  ClientId   string `url:"client_id,omitempty"`
}

func (s *Streams) Channel(channel string) (stream Stream, err error) {
  // uri := fmt.Sprintf("streams/%s", channel)
  // req, err := s.client.NewRequest("GET", uri)
  // if err != nil {
  //   return
  // }

  return
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
