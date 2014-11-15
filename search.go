package twch

type Search struct {
	client *Client
}

type searchChannel struct {
	Channels []Channel   `json:"channels"`
	Total    int         `json:"_total"`
	Links    interface{} `json:"_links"`
}

type searchStream struct {
	Streams []Stream    `json:"streams"`
	Links   interface{} `json:"_links"`
}

type searchGame struct {
	Games []Game      `json:"games"`
	Links interface{} `json:"_links"`
}

type SearchOptions struct {
	Query  string `url:"q"`
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
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

func (s *Searches) Channels(q string, otps *SearchOptions) ([]Channel, err) {
	opts.Query = q
	return nil, nil
}

func (s *Searches) Streams(q string, opts *SearchStreamOptions) ([]Stream, err) {
	opts.Query = q
	return nil, nil
}

func (s *Searches) Games(q string, opts *SearchGameOptions) ([]Game, err) {
	opts.Query = q
	return nil, nil
}
