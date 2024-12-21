package models

type User struct {
	Id        int     `db:"id"`
	Username  *string `db:"username"`
	Name      *string `db:"name"`
	Email     string  `db:"email"`
	Password  string  `db:"password"`
	Bio       *string `db:"bio"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}

type NewUserPayload struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

type LoginPayload struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}
