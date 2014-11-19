package twch

type Ingests struct {
	client *Client
}

type Ingest struct {
	Name         *string  `json:"name,omitempty"`
	Default      *bool    `json:"default,omitempty"`
	ID           *int     `json:"_id,omitempty"`
	URLTemplate  *string  `json:"url_template,omitempty"`
	Availability *float32 `json:"availability,omitempty"`
}

type ingetsResponse struct {
	Ingests []Ingest `json:"ingests,omitempty"`
}

type listIngest struct {
	Ingests []Ingest    `json:"ingests"`
	Links   interface{} `json:"_links"`
}

func (i *Ingests) ListIngests() (ingests []Ingest, resp *Response, err error) {
	req, err := i.client.NewRequest("GET", "ingests")
	if err != nil {
		return
	}

	r := new(ingetsResponse)
	resp, err = i.client.Do(req, r)
	if err != nil {
		return
	}

	ingests = r.Ingests

	return
}
