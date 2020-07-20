package model

type StudyGroup struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}
