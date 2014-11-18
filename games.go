package twch

type Games struct {
	client *Client
}

type Game struct {
	Name        *string `json:"name" `
	Box         *Asset  `json:"box"`
	Logo        *Asset  `json:"logo"`
	GiantbombId *int    `json:"giantbomb_id"`
	Viewers     *int
	Channels    *int
}

type gameRoot struct {
	Top []gameResult `json:"top"`
	listTotal
	listLinks
}

type gameResult struct {
	Game     *Game `json:"game"`
	Viewers  *int  `json:"viewers"`
	Channels *int  `json:"channels"`
}

// Top lists games sorted by number of current viewers on Twitch, most popular first
func (g *Games) ListTop(opts *RequestOptions) (games []Game, resp *Response, err error) {
	url, err := appendOptions("games/top", opts)
	if err != nil {
		return
	}

	req, err := g.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	r := new(gameRoot)
	resp, err = g.client.Do(req, r)
	if err != nil {
		return
	}

	for _, v := range r.Top {
		v.Game.Viewers = v.Viewers
		v.Game.Channels = v.Channels
		games = append(games, *v.Game)
	}

	return games, resp, err
}
