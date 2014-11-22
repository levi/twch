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
	ID        *int    `json:"_id"`
	UpdatedAt *string `json:"updated_at"`
	User      *User   `json:"user"`
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

// AddBlock adds a block to the passed authenticated user. `user` is the current user,
// `target` is the account to block. A successful block returns the new block object.
// This method requires OAuth authentication with the required `user_blocks_edit` scope
func (b *Blocks) AddBlock(user, target string) (block *Block, resp *Response, err error) {
	url := fmt.Sprintf("users/%s/blocks/%s", user, target)
	req, err := b.client.NewRequest("PUT", url)
	if err != nil {
		return
	}

	block = new(Block)
	resp, err = b.client.Do(req, block)
	if err != nil {
		return
	}
	return
}

// RemoveBlock deletes a block from the passed authenticated user. `user` is the current user,
// `target` is the account to block. A 404 error will be returned if the block did not exist
// for the given user.
func (b *Blocks) RemoveBlock(user, target string) (err error) {
	url := fmt.Sprintf("users/%s/blocks/%s", user, target)
	req, err := b.client.NewRequest("DELETE", url)
	if err != nil {
		return
	}

	_, err = b.client.Do(req, nil)
	if err != nil {
		return
	}

	return
}
