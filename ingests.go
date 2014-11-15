package twch

type Ingests struct {
	client *Client
}

type Ingest struct {
  Name         string  `json:"name,omitempty"`
  Default      bool    `json:"default,omitempty"`
  Id           int     `json:"_id,omitempty"`
  UrlTemplate  string  `json:"url_template,omitempty"`
  Availability float32 `json:"availability,omitempty"`
}

type listIngest struct {
  Ingests []Ingest    `json:"ingests"`
  Links   interface{} `json:"_links"`
}

func (i *Ingests) ListIngests() ([]Ingest, error) {
  return nil, nil
}
