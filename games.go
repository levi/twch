package twch

import (
  "net/http"
)

type Games struct {
  client *Client
}

type Asset struct {
  Large    string `json:"large"`
  Medium   string `json:"medium"`
  Small    string `json:"small"`
  Template string `json:"template"`
}

type Game struct {
  Name        string `json:"name" `
  Box         *Asset `json:"box"`
  Logo        *Asset `json:"logo"`
  GiantbombId int    `json:"giantbomb_id"`
  Viewers     int
  Channels    int
}

type gameRoot struct {
  Links interface{}  `json:"_links"`
  Total int          `json:"_total"`
  Top   []gameResult `json:"top"`
}

type gameResult struct {
  Game     *Game `json:"game"`
  Viewers  int   `json:"viewers"`
  Channels int   `json:"channels"`
}

// Top lists games sorted by number of current viewers on Twitch, most popular first
func (g *Games) Top() (games []Game, resp *http.Response, err error) {
  req, err := g.client.NewRequest("GET", "games/top")
  if err != nil {
    return
  }

  res := new(gameRoot)
  resp, err = g.client.Do(req, res)
  if err != nil {
    return
  }

  for _, v := range res.Top {
    v.Game.Viewers = v.Viewers
    v.Game.Channels = v.Channels
    games = append(games, *v.Game)
  }

  return games, resp, err
}
