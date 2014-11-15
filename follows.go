package twch

type Follows struct {
	client *Client
}

type Follow struct {
  User      User   `json:"user,omitempty"`
  CreatedAt string `json:"created_at,omitempty"`
}

type FollowOptions struct {
  Direction string `url:"direction,omitempty"`
  RequestOptions
}

func (f *Follows) GetChannelFollows(channel string) ([]Follow, error) {
  // "channels/:channel/follows"
  return nil, nil
}

func (f *Follows) GetFollowedChannels(user string) ([]Channel, error) {
  // "users/:user/follows/channels"
  return nil, nil
}

func (f *Follows) IsFollowing(user, channel string) (Channel, error) {
  // "users/:user/follows/channels/:target"
  return Channel{}, nil
}

func (f *Follows) FollowChannel(user, channel string) (Channel, error) {
  // PUT "users/:user/follows/channels/:target"
  return Channel{}, nil
}

func (f *Follows) UnfollowChannel(user, channel string) error {
  // DELETE "users/:user/follows/channels/:target"
  return nil
}
