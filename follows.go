package twch

type Follow struct {
	User      User   `json:"user,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type FollowOptions struct {
	Direction string `url:"direction,omitempty"`
	RequestOptions
}
