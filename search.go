package twch

type Search struct {
	client *Client
}

type searchChannels struct {
	Channels []Channel `json:"channels"`
	listTotal
	listLinks
}

type searchStreams struct {
	Streams []Stream `json:"streams"`
	listTotal
	listLinks
}

type searchGames struct {
	Games []Game `json:"games"`
}

type searchChannelOptions struct {
	Query string `url:"q"`
	ListOptions
}

type searchStreamOptions struct {
	Query string `url:"q"`
	RequestOptions
}

type searchGameOptions struct {
	Type  string `url:"type"`
	Live  bool   `url:"live,omitempty"`
	Query string `url:"q"`
}

// Channels returns a list of channels matching the search query
func (s *Search) Channels(q string, opts *ListOptions) (ch []Channel, resp *Response, err error) {
	o := &searchChannelOptions{Query: q}
	if opts != nil {
		o.ListOptions = *opts
	}

	url, err := appendOptions("search/channels", o)
	if err != nil {
		return
	}

	req, err := s.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	r := new(searchChannels)
	resp, err = s.client.Do(req, r)
	if err != nil {
		return
	}

	ch = r.Channels

	return
}

// Streams returns a list of streams matching the search query
func (s *Search) Streams(q string, opts *RequestOptions) (st []Stream, resp *Response, err error) {
	o := &searchStreamOptions{Query: q}
	if opts != nil {
		o.RequestOptions = *opts
	}

	url, err := appendOptions("search/streams", o)
	if err != nil {
		return
	}

	req, err := s.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	r := new(searchStreams)
	resp, err = s.client.Do(req, r)
	if err != nil {
		return
	}

	st = r.Streams

	return
}

// Games returns a list of games matching the search query
func (s *Search) Games(q string, live bool) (g []Game, resp *Response, err error) {
	o := &searchGameOptions{Query: q, Type: "suggest"}
	if live {
		o.Live = true
	}

	url, err := appendOptions("search/games", o)
	if err != nil {
		return
	}

	req, err := s.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	r := new(searchGames)
	resp, err = s.client.Do(req, r)
	if err != nil {
		return
	}

	g = r.Games

	return
}
