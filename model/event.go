package model

import (
	"time"
)

type Event struct {
	ID           uint64    `db:"id"`
	StudyGroupID uint64    `db:"study_group_id"`
	Name         string    `db:"name"`
	EventDate    time.Time `db:"event_date"`
	Description  string    `db:"description"`
	Place        string    `db:"place"`
}
