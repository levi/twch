package twch

type Channels struct {
	client *Client
}

type Channel struct {
  Id        int    `json:"_id,omitempty"`
  Name      string `json:"name,omitempty"`
  Game      string `json:"game,omitempty"`
  CreatedAt string `json:"created_at,omitempty"`
  Delay     int    `json:"delay,omitempty"`
  StreamKey string `json:"stream_key,omitempty"`
  Teams     []struct {
    Id          int         `json:"_id,omitempty"`
    Name        string      `json:"name,omitempty"`
    CreatedAt   string      `json:"created_at,omitempty"`
    UpdatedAt   string      `json:"updated_at,omitempty"`
    Background  string      `json:"background,omitempty"`
    Banner      string      `json:"banner,omitempty"`
    Logo        string      `json:"logo,omitempty"`
    Info        string      `json:"info,omitempty"`
    DisplayName string      `json:"display_name,omitempty"`
    Links       interface{} `json:"_links,omitempty"`
  } `json:"teams,omitempty"`
  Status      string      `json:"status,omitempty"`
  UpdatedAt   string      `json:"updated_at,omitempty"`
  Banner      string      `json:"banner,omitempty"`
  VideoBanner string      `json:"video_banner,omitempty"`
  Background  string      `json:"background,omitempty"`
  Logo        string      `json:"logo,omitempty"`
  Url         string      `json:"url,omitempty"`
  Login       string      `json:"login,omitempty"`
  DisplayName string      `json:"display_name,omitempty"`
  Email       string      `json:"email,omitempty"`
  Mature      bool        `json:"mature,omitempty"`
  Links       interface{} `json:"_links,omitempty"`
}

type ChannelOptions struct {
  Status string `url:"status,omitempty"`
  Game   string `url:"game,omitempty"`
  Delay  string `url:"deplay,omitempty"`
}

type CommericalOptions struct {
  Length string `url:"length,omitempty"`
}

func (c *Channels) GetChannel(channel string) (Channel, error) {
  // "channels/:channel"
  return Channel{}, nil
}

func (c *Channels) GetUserChannel() (Channel, error) {
  // "channel"
  return Channel{}, nil
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
