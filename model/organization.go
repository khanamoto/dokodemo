package model

type Organization struct {
	ID   uint64 `db:"id"`
	Name string `db:"name" validate:"required,max=255"`
	URL  string `db:"url" validate:"required,alphanum,max=255"`
}
