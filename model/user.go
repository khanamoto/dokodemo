package model

type User struct {
	ID        uint64 `db:"id"`
	Name      string `db:"name"`
	UserName  string `db:"user_name"`
	Email     string `db:"email"`
	Job       string `db:"job"`
	WebSite   string `db:"website"`
	Biography string `db:"biography"`
}
