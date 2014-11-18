package twch

type Users struct {
	client *Client
}

type User struct {
	ID          *int    `json:"_id,omitempty"`
	Type        *string `json:"type,omitempty"`
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	Bio         *string `json:"bio,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	Email       *string `json:"email,omitempty"`
	Partnered   *bool   `json:"partnered,omitempty"`
	Staff       *bool   `json:"staff,omitempty"`
}

func (u *Users) CurrentUser() (User, error) {
	return User{}, nil
}

func (u *Users) User(username string) (User, error) {
	return User{}, nil
}
