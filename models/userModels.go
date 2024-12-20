package models

type User struct {
	Id        int     `json:"id"`
	Username  string  `json:"username"`
	Name      *string `json:"name"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Bio       *string `json:"bio"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type NewUserPayload struct {
	Username string  `json:"username"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Bio      *string `json:"bio"`
}
