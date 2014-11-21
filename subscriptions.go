package twch

import (
	"fmt"
)

type Subscriptions struct {
	client *Client
}

type listSubscription struct {
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
	*listLinks
	*listTotal
}

type Subscription struct {
	ID        *string  `json:"_id,omitempty"`
	User      *User    `json:"user,omitempty"`
	Channel   *Channel `json:"channel,omitempty"`
	CreatedAt *string  `json:"created_at,omitempty"`
}

type SubscriptionOptions struct {
	Direction string `url:"direction,omitempty"`
	ListOptions
}

// GetChannelSubscriptions returns a list of subscriptions for the given channel,
// ordered by creation date.
// Requires the `channel_subscriptions` authentication scope.
func (s *Subscriptions) GetChannelSubscriptions(channel string, opts *SubscriptionOptions) (sub []Subscription, resp *Response, err error) {
	url := fmt.Sprintf("channels/%s/subscriptions", channel)
	url, err = appendOptions(url, opts)
	if err != nil {
		return
	}

	req, err := s.client.NewRequest("GET", url, false)
	if err != nil {
		return
	}

	r := new(listSubscription)
	resp, err = s.client.Do(req, r)
	if err != nil {
		return
	}

	sub = r.Subscriptions

	return
}

// UserSubscribed returns a subscription if the user is subscribed to
// the given channel.
// Requires the `channel_check_subscription` authentication scope for
// the given channel.
func (s *Subscriptions) GetUserSubscribed(channel, user string) (sub *Subscription, resp *Response, err error) {
	url := fmt.Sprintf("channels/%s/subscriptions/%s", channel, user)
	req, err := s.client.NewRequest("GET", url, false)
	if err != nil {
		return
	}

	sub = new(Subscription)
	resp, err = s.client.Do(req, sub)
	if err != nil {
		return
	}

	return
}

// ChannelSubscribed returns a subscription with a channel that the user
// subscribes to.
// Requires the `user_subscriptions` authentication scope for the given user.
func (s *Subscriptions) GetSubscribedChannel(user, channel string) (sub *Subscription, resp *Response, err error) {
	url := fmt.Sprintf("users/%s/subscriptions/%s", user, channel)
	req, err := s.client.NewRequest("GET", url, false)
	if err != nil {
		return
	}

	sub = new(Subscription)
	resp, err = s.client.Do(req, sub)
	if err != nil {
		return
	}

	return
}
