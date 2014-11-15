package twch

type Root struct {
	Link struct {
		Channel  string `json:"channel"`
		Users    string `json:"users"`
		User     string `json:"user"`
		Channels string `json:"channels"`
		Chat     string `json:"chat"`
		Streams  string `json:"streams"`
		Ingests  string `json:"ingests"`
	} `json:"_links"`

	Token struct {
		Valid         bool   `json:"valid"`
		UserName      string `json:"user_name"`
		Authorization struct {
			Scopes    []string `json:"scopes"`
			CreatedAt string   `json:"created_at"`
			UpdatedAt string   `josn:"updated_at"`
		}
	} `json:"token"`
}
