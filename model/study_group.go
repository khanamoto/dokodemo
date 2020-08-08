package model

type StudyGroup struct {
	ID           uint64 `db:"id"`
	DepartmentID uint64 `db:"department_id"`
	Name         string `db:"name" validate:"required,max=255"`
	URL          string `db:"url" validate:"required,alphanum,max=255"`
}
