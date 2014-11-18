package twch

import (
	"fmt"
)

type Blocks struct {
	client *Client
}

type listBlocks struct {
	Blocks []Block `json:"blocks"`
	listLinks
}

type Block struct {
	ID        int    `json:"_id"`
	UpdatedAt string `json:"updated_at"`
	User      User   `json:"user"`
}

func (b *Blocks) ListBlocks(login string, opts *ListOptions) (blocks []Block, resp *Response, err error) {
	url := fmt.Sprintf("users/%s/blocks", login)
	u, err := appendOptions(url, opts)
	if err != nil {
		return
	}

	req, err := b.client.NewRequest("GET", u)
	if err != nil {
		return
	}

	r := new(listBlocks)
	resp, err = b.client.Do(req, r)
	if err != nil {
		return
	}
	blocks = r.Blocks
	return
}

func (b *Blocks) BlockUser(target, current string) (Block, error) {
	// PUT "users/:user/blocks/:target"
	return Block{}, nil
}

func (b *Blocks) UnblockUser(target, current string) error {
	// DELETE "users/:users/blocks/:target"
	return nil
}
