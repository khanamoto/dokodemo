package model

type SubMembership struct {
	ID              uint64 `db:"id"`
	UserID          uint64 `db:"user_id"`
	SubStudyGroupID uint64 `db:"sub_study_group_id"`
	Authority       int32  `db:"authority" validate:"required,oneof=1 2"`
}
