package twch

type Search struct {
	client *Client
}

type searchChannel struct {
	Channels []Channel `json:"channels"`
	listTotal
	listLinks
}

type searchStream struct {
	Streams []Stream `json:"streams"`
	listLinks
}

type searchGame struct {
	Games []Game `json:"games"`
	listLinks
}

type SearchOptions struct {
	Query string `url:"q"`
	ListOptions
}

type SearchStreamOptions struct {
	HLS bool `url:"hls,omitempty"`
	SearchOptions
}

type SearchGameOptions struct {
	Type string `url:"url"`
	Live bool   `url:"live,omitempty"`
	SearchOptions
}

// Channels returns a list of channels matching the search query
func (s *Search) Channels(q string, opts *SearchOptions) (ch []Channel, resp *Response, err error) {
	opts.Query = q
	url, err := appendOptions("search/channels", opts)
	if err != nil {
		return
	}

	req, err := s.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	r := new(searchChannel)
	resp, err = s.client.Do(req, r)
	if err != nil {
		return
	}

	ch = r.Channels

	return
}

func (s *Search) Streams(q string, opts *SearchStreamOptions) ([]Stream, error) {
	opts.Query = q
	return nil, nil
}

func (s *Search) Games(q string, opts *SearchGameOptions) ([]Game, error) {
	opts.Query = q
	return nil, nil
}
