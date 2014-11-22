package twch

type Chats struct {
	client *Client
}

type listEmoticons struct {
	Emoticons []Emoticon  `json:"emoticons,omitempty"`
	Links     interface{} `json:"_links,omitempty"`
}

type EmoticonImage struct {
	EmoticonSet *int    `json:"emoticon_set,omitempty"`
	Height      *int    `json:"height,omitempty"`
	Width       *int    `json:"width,omitempty"`
	URL         *string `json:"url,omitempty"`
}

type Emoticon struct {
	Regex  *string         `json:"regex"`
	Images []EmoticonImage `json:"images"`
}

// ListEmoticons returns a list of all the emoticons on Twitch
func (c *Chats) ListEmoticons() (e []Emoticon, resp *Response, err error) {
	req, err := c.client.NewRequest("GET", "chat/emoticons")
	if err != nil {
		return
	}

	r := new(listEmoticons)
	resp, err = c.client.Do(req, r)
	if err != nil {
		return
	}

	e = r.Emoticons

	return
}
