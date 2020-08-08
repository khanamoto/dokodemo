package model

type Belonging struct {
	ID             uint64 `db:"id"`
	UserID         uint64 `db:"user_id"`
	OrganizationID uint64 `db:"organization_id"`
	Authority      int32  `db:"authority" validate:"required,oneof=1 2"`
}
