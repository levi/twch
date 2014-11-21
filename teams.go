package twch

import (
	"fmt"
)

type Teams struct {
	client *Client
}

type Team struct {
	ID          *int    `json:"_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Info        *string `json:"info,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	Banner      *string `json:"banner,omitempty"`
	Background  *string `json:"background,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

type teamsResponse struct {
	Teams []Team `json:"teams,omitempty"`
	*listLinks
}

// ListTeams returns a list of teams that are active on Twitch
func (t *Teams) ListTeams() (teams []Team, resp *Response, err error) {
	req, err := t.client.NewRequest("GET", "teams", false)
	if err != nil {
		return
	}

	r := new(teamsResponse)
	resp, err = t.client.Do(req, r)
	if err != nil {
		return
	}

	teams = r.Teams

	return
}

// GetTeam returns a team for the passed team name
func (t *Teams) GetTeam(team string) (r *Team, resp *Response, err error) {
	url := fmt.Sprintf("teams/%s", team)
	req, err := t.client.NewRequest("GET", url, false)
	if err != nil {
		return
	}

	r = new(Team)
	resp, err = t.client.Do(req, r)
	if err != nil {
		return
	}

	return
}
