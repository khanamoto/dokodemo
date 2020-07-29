package model

type User struct {
	ID        uint64 `db:"id"`
	Name      string `db:"name" validate:"required,min=1,max=32"`
	UserName  string `db:"user_name" validate:"required,alphanum,min=4,max=32"`
	Email     string `db:"email" validate:"required,email,max=255"`
	Password  string `db:"password" validate:"required,min=6"`
	Job       string `db:"job" validate:"max=255"`
	WebSite   string `db:"website" validate:"url,max=255"`
	Biography string `db:"biography" validate:"max=1000"`
}
