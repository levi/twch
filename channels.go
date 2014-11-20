package twch

import (
	"fmt"
)

type Channels struct {
	client *Client
}

type editorsResponse struct {
	Users []User `json:"users,omitempty"`
}

type Channel struct {
	ID                           *int    `json:"_id,omitempty"`
	DisplayName                  *string `json:"display_name,omitempty"`
	Name                         *string `json:"name,omitempty"`
	Title                        *string `json:"title,omitempty"`
	Game                         *string `json:"game,omitempty"`
	Delay                        *int    `json:"delay,omitempty"`
	StreamKey                    *string `json:"stream_key,omitempty"`
	Teams                        []Team  `json:"teams,omitempty"`
	Status                       *string `json:"status,omitempty"`
	Banner                       *string `json:"banner,omitempty"`
	ProfileBanner                *string `json:"profile_banner,omitempty"`
	ProfileBannerBackgroundColor *string `json:"profile_banner_background_color,omitempty"`
	VideoBanner                  *string `json:"video_banner,omitempty"`
	Background                   *string `json:"background,omitempty"`
	Logo                         *string `json:"logo,omitempty"`
	URL                          *string `json:"url,omitempty"`
	Login                        *string `json:"login,omitempty"`
	Email                        *string `json:"email,omitempty"`
	Mature                       *bool   `json:"mature,omitempty"`
	Language                     *string `json:"language,omitempty"`
	BroadcasterLanguage          *string `json:"broadcaster_language,omitempty"`
	Partner                      *bool   `json:"partner,omitempty"`
	Views                        *int    `json:"views,omitempty"`
	Followers                    *int    `json:"followers,omitempty"`
	CreatedAt                    *string `json:"created_at,omitempty"`
	UpdatedAt                    *string `json:"updated_at,omitempty"`
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

// ListChannelEditors returns a list of user objects associated with the channel
// as "editor" status.
// This method requires the `channel_read` authentication scope
func (c *Channels) ListChannelEditors(channel string) (u []User, resp *Response, err error) {
	url := fmt.Sprintf("channels/%s/editors", channel)
	req, err := c.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	e := new(editorsResponse)
	resp, err = c.client.Do(req, e)
	if err != nil {
		return
	}

	u = e.Users

	return
}

// ListChannelTeams returns a list of teams for the given channel
func (c *Channels) ListChannelTeams(channel string) (t []Team, resp *Response, err error) {
	url := fmt.Sprintf("channels/%s/teams", channel)
	req, err := c.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	e := new(teamsResponse)
	resp, err = c.client.Do(req, e)
	if err != nil {
		return
	}

	t = e.Teams

	return
}

func (c *Channels) UpdateChannel(channel string) error {
	// PUT "channels/:channel"
	return nil
}

// ResetStreamKey reset's an authenticated channel's stream key
// Requires the `channel_stream` authentication scope
func (c *Channels) ResetStreamKey(channel string) (err error) {
	url := fmt.Sprintf("channels/%s/stream_key", channel)
	req, err := c.client.NewRequest("DELETE", url)
	if err != nil {
		return
	}

	_, err = c.client.Do(req, nil)
	if err != nil {
		return
	}

	return
}

func (c *Channels) StartCommercial(channel string) error {
	// POST "channels/:channel/commerical"
	return nil
}
