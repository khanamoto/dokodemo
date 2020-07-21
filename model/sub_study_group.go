package model

type SubStudyGroup struct {
	ID           uint64 `db:"id"`
	StudyGroupID uint64 `db:"study_group_id"`
	Name         string `db:"name"`
}
