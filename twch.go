package twch

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "net/url"
)

const (
  baseUrl      = "https://api.twitch.tv/kraken/"
  acceptHeader = "application/vnd.twitchtv.v3+json"
)

type Client struct {
  client *http.Client

  id string

  BaseUrl *url.URL

  Blocks        *Blocks
  Channels      *Channels
  Chat          *Chats
  Follows       *Follows
  Games         *Games
  Ingests       *Ingests
  Search        *Search
  Streams       *Streams
  Subscriptions *Subscriptions
  Teams         *Teams
  Users         *Users
  Videos        *Videos
}

// NewClient constructs a new client to interface with the Twitch API
func NewClient(id string) (client *Client, err error) {
  client = new(Client)
  client.id = id
  client.client = http.DefaultClient
  client.BaseUrl, _ = url.Parse(baseUrl)
  client.Blocks = &Blocks{client: client}
  client.Channels = &Channels{client: client}
  client.Chat = &Chats{client: client}
  client.Follows = &Follows{client: client}
  client.Games = &Games{client: client}
  client.Ingests = &Ingests{client: client}
  client.Search = &Search{client: client}
  client.Streams = &Streams{client: client}
  client.Subscriptions = &Subscriptions{client: client}
  client.Teams = &Teams{client: client}
  client.Users = &Users{client: client}
  client.Videos = &Videos{client: client}
  return client, nil
}

// NewRequest constructs a valid http.Request object for Twitch requests
func (c *Client) NewRequest(method, uri string) (req *http.Request, err error) {
  apiUri, err := url.Parse(uri)
  if err != nil {
    return nil, err
  }

  reqUrl := c.BaseUrl.ResolveReference(apiUri)
  q, err := url.ParseQuery(reqUrl.RawQuery)
  if err != nil {
    return nil, err
  }
  q.Add("client_id", c.id)
  reqUrl.RawQuery = q.Encode()

  req, err = http.NewRequest(method, reqUrl.String(), nil)
  if err != nil {
    return nil, err
  }

  req.Header.Add("Accept", acceptHeader)

  return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (resp *http.Response, err error) {
  resp, err = c.client.Do(req)
  if err != nil {
    return resp, err
  }
  defer resp.Body.Close()

  err = json.NewDecoder(resp.Body).Decode(v)
  if err != nil {
    return resp, err
  }

  return resp, nil
}

func main() {
  c, err := NewClient("test-client-id")
  if err != nil {
    log.Fatal(err)
  }

  games, _, err := c.Games.Top()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%+v\n", games)

  viewers, channels, err := c.Streams.Summary()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Viewers: %d, Channels: %d", viewers, channels)
}
