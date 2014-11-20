package twch

import (
	"fmt"
)

type Videos struct {
	client *Client
}

type videosResponse struct {
	Videos []Video     `json:"videos"`
	Links  interface{} `json:"_links"`
}

type Video struct {
	ID          *string  `json:"_id"`
	Title       *string  `json:"title"`
	URL         *string  `json:"url"`
	Views       *int     `json:"views"`
	Description *string  `json:"description"`
	Length      *int     `json:"length"`
	Game        *string  `json:"game"`
	Preview     *string  `json:"preview"`
	RecordedAt  *string  `json:"recorded_at"`
	Channel     *Channel `json:"channel"`
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

// GetVideo returns a video object via its ID
func (v *Videos) GetVideo(id int) (video *Video, resp *Response, err error) {
	url := fmt.Sprintf("videos/%d", id)
	req, err := v.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	video = new(Video)
	resp, err = v.client.Do(req, video)
	if err != nil {
		return
	}

	return
}

func (v *Videos) ListTop() ([]Video, error) {
	// "videos/top"
	return nil, nil
}

func (v *Videos) ListChannelVideos(channel string) ([]Video, error) {
	// "channels/:channel/videos"
	return nil, nil
}
