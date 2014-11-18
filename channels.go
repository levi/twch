package twch

import (
	"fmt"
)

type Channels struct {
	client *Client
}

type Channel struct {
	ID          *int    `json:"_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Name        *string `json:"name,omitempty"`
	Title       *string `json:"title,omitempty"`
	Game        *string `json:"game,omitempty"`
	Delay       *int    `json:"delay,omitempty"`
	StreamKey   *string `json:"stream_key,omitempty"`
	Teams       []Team  `json:"teams,omitempty"`
	Status      *string `json:"status,omitempty"`
	Banner      *string `json:"banner,omitempty"`
	VideoBanner *string `json:"video_banner,omitempty"`
	Background  *string `json:"background,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	URL         *string `json:"url,omitempty"`
	Login       *string `json:"login,omitempty"`
	Email       *string `json:"email,omitempty"`
	Mature      *bool   `json:"mature,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

type ChannelOptions struct {
	Status string `url:"status,omitempty"`
	Game   string `url:"game,omitempty"`
	Delay  string `url:"deplay,omitempty"`
}

type CommericalOptions struct {
	Length string `url:"length,omitempty"`
}

// GetChannel returns a channel by name
func (c *Channels) GetChannel(channel string) (ch *Channel, resp *Response, err error) {
	url := fmt.Sprintf("channels/%s", channel)
	req, err := c.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	ch = new(Channel)
	resp, err = c.client.Do(req, ch)
	if err != nil {
		return
	}

	return
}

// GetUserChannel returns the channel for the authenticated user
// Requires the `channel_read` authentication scope to be approved
func (c *Channels) GetUserChannel() (ch *Channel, resp *Response, err error) {
	req, err := c.client.NewRequest("GET", "channel")
	if err != nil {
		return
	}

	ch = new(Channel)
	resp, err = c.client.Do(req, ch)
	if err != nil {
		return
	}

	return
}

func (c *Channels) GetEditors(channel string) ([]User, error) {
	// "channels/:channel/editors"
	return nil, nil
}

func (c *Channels) GetTeams(channel string) ([]Team, error) {
	// "channels/:channel/teams"
	return nil, nil
}

func (c *Channels) UpdateChannel(channel string) error {
	// PUT "channels/:channel"
	return nil
}

func (c *Channels) ResetStreamKey(channel string) error {
	// DELETE "channels/:channel/stream_key"
	return nil
}

func (c *Channels) StartCommercial(channel string) error {
	// POST "channels/:channel/commerical"
	return nil
}
