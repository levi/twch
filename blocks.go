package twch

type Blocks struct {
	client *Client
}

type listBlocks struct {
  Blocks []Block     `json:"blocks"`
  Links  interface{} `json:"_links"`
}

type Block struct {
  Id        int    `json:"_id"`
  UpdatedAt string `json:"updated_at"`
  User      User   `json:"user"`
}

func (b *Blocks) ListBlocks(login string) ([]Block, error) {
  // "users/:login/blocks"
  return nil, nil
}

func (b *Blocks) BlockUser(target, current string) (Block, error) {
  // PUT "users/:user/blocks/:target"
  return Block{}, nil
}

func (b *Blocks) UnblockUser(target, current string) error {
  // DELETE "users/:users/blocks/:target"
  return nil
}
