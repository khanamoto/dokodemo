package model

type Membership struct {
	ID           uint64 `db:"id"`
	UserID       uint64 `db:"user_id"`
	StudyGroupID uint64 `db:"study_group_id"`
	Authority    int32  `db:"authority" validate:"required,oneof=1 2"`
}
