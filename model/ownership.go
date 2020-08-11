package model

type Ownership struct {
	ID           uint64 `db:"id"`
	StudyGroupID uint64 `db:"study_group_id"`
	EventID      uint64 `db:"event_id"`
}
