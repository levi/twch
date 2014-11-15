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

func (f *Follows) GetChannelFollows(channel string) ([]Follow, err) {
  // "channels/:channel/follows"
  return nil, nil
}

func (f *Follows) GetFollowedChannels(user string) ([]Channel, err) {
  // "users/:user/follows/channels"
  return nil, nil
}

func (f *Follows) IsFollowing(user, channel string) (Channel, err) {
  // "users/:user/follows/channels/:target"
  return nil, nil
}

func (f *Follows) FollowChannel(user, channel string) (Channel, err) {
  // PUT "users/:user/follows/channels/:target"
  return nil, nil
}

func (f *Follows) UnfollowChannel(user, channel string) err {
  // DELETE "users/:user/follows/channels/:target"
  return nil, nil
}
