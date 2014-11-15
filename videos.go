package twch

type Videos struct {
	client *Client
}

type videosResponse struct {
  Videos []video     `json:"videos"`
  Links  interface{} `json:"_links"`
}

type Video struct {
  Id          string `json:"_id"`
  Title       string `json:"title"`
  Url         string `json:"url"`
  Views       int    `json:"views"`
  Description string `json:"description"`
  Length      int    `json:"length"`
  Game        string `json:"game"`
  Preview     string `json:"preview"`
  RecordedAt  string `json:"recorded_at"`
  Channel     struct {
    Name        string `json:"name"`
    DisplayName string `json:"display_name"`
  } `json:"channel"`
  Links interface{} `json:"_links"`
}

type VideoListRequestOptions struct {
  RequestOptions
  Game   string
  Period string
}

type VideoChannelOptions struct {
  limit      int
  offset     int
  broadcasts bool
}

func (v *Videos) Video(id int) (Video, err) {
  // "videos/:id"
  return nil, nil
}

func (v *Videos) Top() ([]Video, err) {
  // "videos/top"
  return nil, nil
}

func (v *Videos) ChannelVideos(channel string) ([]Video, err) {
  // "channels/:channel/videos"
  return nil, nil
}
