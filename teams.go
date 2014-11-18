package twch

type Teams struct {
	client *Client
}

type Team struct {
	ID          *int    `json:"_id"`
	Name        *string `json:"name"`
	Info        *string `json:"info"`
	DisplayName *string `json:"display_name"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
	Logo        *string `json:"logo"`
	Banner      *string `json:"banner"`
	Background  *string `json:"background"`
}

func (t *Teams) ListTeams() ([]Team, error) {
	// "teams"
	return nil, nil
}

func (t *Teams) GetTeam(team string) (Team, error) {
	// "teams/:team"
	return Team{}, nil
}
