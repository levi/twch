package twch

type Chats struct {
	client *Client
}

type listEmoticons struct {
	Emoticons []Emoticon  `json:"emoticons"`
	Links     interface{} `json:"_links"`
}

type EmoticonImage struct {
	EmoticonSet int `json:"emoticon_set"`
	Height      int `json:"height"`
	Width       int `json:"width"`
	Url         int `json:"url"`
}

type Emoticon struct {
	Regex  string          `json:"regex"`
	Images []EmoticonImage `json:"images"`
}

func (c *Chats) GetEmoticons() ([]Emoticon, error) {
	// "chat/emoticons"
	return nil, nil
}
