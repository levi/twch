package twch

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

// GetCurrentUser returns
func (u *Users) GetCurrentUser() (User, error) {
	return User{}, nil
}

func (u *Users) GetUser(username string) (User, error) {
	return User{}, nil
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
