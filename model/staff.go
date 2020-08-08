package model

type Staff struct {
	ID           uint64 `db:"id"`
	UserID       uint64 `db:"user_id"`
	DepartmentID uint64 `db:"department_id"`
	Authority    int32  `db:"authority" validate:"required,oneof=1 2"`
}
