package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var membershipNotFoundError = model.NotFoundError("membership")

func (r *repository) CreateMembership(userID uint64, studyGroupID uint64, authority int32) (*model.Membership, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO membership (id, user_id, study_group_id, authority, created_at, updated_at) VALUES (?,?,?,?,?,?)`,
		id, userID, studyGroupID, authority, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Membership{ID: id, UserID: userID, StudyGroupID: studyGroupID, Authority: authority}, nil
}
