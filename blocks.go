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

func (b *Blocks) ListBlocks(login string) ([]Block, err) {
  // "users/:login/blocks"
  return nil, nil
}

func (b *Blocks) BlockUser(target, current string) (Block, err) {
  // PUT "users/:user/blocks/:target"
  return nil
}

func (b *Blocks) UnblockUser(target, current string) err {
  // DELETE "users/:users/blocks/:target"
  return nil
}
