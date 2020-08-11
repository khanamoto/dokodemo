package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var ownershipNotFoundError = model.NotFoundError("ownership")

func (r *repository) CreateOwnership(studyGroupID uint64, eventID uint64) (*model.Ownership, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO ownership (id, study_group_id, event_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		id, studyGroupID, eventID, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Ownership{ID: id, StudyGroupID: studyGroupID, EventID: eventID}, nil
}
