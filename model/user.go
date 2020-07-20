package model

type User struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	UserName string `db:"user_name"`
	Email    string `db:"email"`
}
