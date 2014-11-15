package twch

type Subscriptions struct {
	client *Client
}

type listSubscription struct {
  Total int `json:"_total,omitempty"`
  Links struct {
    Next string `json:"next"`
    Self string `json:"self"`
  } `json:"_links,omitempty"`
  Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

type Subscription struct {
  Id        int    `json:"_id,omitempty"`
  User      User   `json:"user,omitempty"`
  CreatedAt string `json:"created_at,omitempty"`
}

type SubscriptionOptions struct {
  Direction string `url:"direction,omitempty"`
}

func (s *Subscriptions) GetChannelSubscriptions(channel string, opts *SubscriptionOptions) ([]Subscription, err) {
  // "channels/:channel/subscriptions"
  return nil, nil
}

func (s *Subscriptions) UserSubscribed(user, channel string) (Subscription, err) {
  // "channels/:channel/subscriptions/:user"
  return nil, nil
}

func (s *Subscriptions) ChannelSubscribed(user, channel string) (Subscription, err) {
  // "users/:user/subscriptions/:channel"
  return nil, nil
}
