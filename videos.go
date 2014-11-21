package twch

import (
	"fmt"
)

type Videos struct {
	client *Client
}

type videosResponse struct {
	Videos []Video `json:"videos"`
	*listLinks
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

type VideoRequestOptions struct {
	Game   string `url:"game,omitempty"`
	Period string `url:"period,omitempty"`
	ListOptions
}

type VideoChannelOptions struct {
	Broadcasts bool `url:"broadcasts,omitempty"`
	ListOptions
}

// GetVideo returns a video object via its ID
func (v *Videos) GetVideo(id int) (video *Video, resp *Response, err error) {
	url := fmt.Sprintf("videos/%d", id)
	req, err := v.client.NewRequest("GET", url, false)
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

// ListTop returns a list of the top videos on twitch for the specified period of time,
// ordered by most popular first. Defined time periods are "week", "month", or "all".
// By default, the top videos of the "week" are returned.
// Videos belonging to a specific game can be returned by passing the name of the game in the
// `Game` VideoRequestOption value. Otherwise, all games will be included in the result.
func (v *Videos) ListTop(opts *VideoRequestOptions) (videos []Video, resp *Response, err error) {
	url, err := appendOptions("videos/top", opts)
	if err != nil {
		return
	}

	req, err := v.client.NewRequest("GET", url, false)
	if err != nil {
		return
	}

	r := new(videosResponse)
	resp, err = v.client.Do(req, r)
	if err != nil {
		return
	}

	videos = r.Videos

	return
}

// ListChannelVideos returns videos belonging to the target channel.
// Only broadcasts will be returned when the `VideoChannelOptions.Broadcasts` field
// is true. Otherwise only highlights are returned by default.
func (v *Videos) ListChannelVideos(channel string, opts *VideoChannelOptions) (videos []Video, resp *Response, err error) {
	url := fmt.Sprintf("channels/%s/videos", channel)
	url, err = appendOptions(url, opts)
	if err != nil {
		return
	}

	req, err := v.client.NewRequest("GET", url, false)
	if err != nil {
		return
	}

	r := new(videosResponse)
	resp, err = v.client.Do(req, r)
	if err != nil {
		return
	}

	videos = r.Videos

	return
}
