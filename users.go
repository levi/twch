package twch

import (
	"fmt"
)

type Users struct {
	client *Client
}

type User struct {
	ID          *int    `json:"_id,omitempty"`
	Type        *string `json:"type,omitempty"`
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	Bio         *string `json:"bio,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	Email       *string `json:"email,omitempty"`
	Partnered   *bool   `json:"partnered,omitempty"`
	Staff       *bool   `json:"staff,omitempty"`
}

// GetUser returns the public profile of a given Twitch user
func (u *Users) GetUser(username string) (user *User, resp *Response, err error) {
	url := fmt.Sprintf("users/%s", username)
	req, err := u.client.NewRequest("GET", url)
	if err != nil {
		return
	}

	user = new(User)
	resp, err = u.client.Do(req, user)
	if err != nil {
		return
	}

	return
}

// GetCurrentUser returns the authenticated user with email and partner info.
// Requires the `user_read` authentication scope.
func (u *Users) GetCurrentUser() (user *User, resp *Response, err error) {
	req, err := u.client.NewRequest("GET", "user")
	if err != nil {
		return
	}

	user = new(User)
	resp, err = u.client.Do(req, user)
	if err != nil {
		return
	}

	return
}

func (u *Users) ListFollowedStreams(opts *RequestOptions) (s []Stream, rep *Response, err error) {
	return nil, nil, nil
}

func (u *Users) ListFollowedVideos(opts *ListOptions) (s []Video, resp *Response, err error) {
	return nil, nil, nil
}

func (u *Users) ListFollowedChannels(user string) (ch []Channel, resp *Response, err error) {
	// "users/:user/follows/channels"
	return nil, nil, nil
}

func (u *Users) IsFollowing(user, channel string) (ch *Channel, resp *Response, err error) {
	// "users/:user/follows/channels/:target"
	return nil, nil, nil
}

func (u *Users) FollowChannel(user, channel string) (ch *Channel, resp *Response, err error) {
	// PUT "users/:user/follows/channels/:target"
	return nil, nil, nil
}

func (u *Users) UnfollowChannel(user, channel string) (err error) {
	// DELETE "users/:user/follows/channels/:target"
	return nil
}
