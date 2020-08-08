package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var membershipNotFoundError = model.NotFoundError("membership")

func (r *repository) CreateMembership(userIDs []uint64, studyGroupID uint64, authority int32) (*model.Membership, error) {
	now := time.Now()

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		`INSERT INTO membership (id, user_id, study_group_id, authority, created_at, updated_at) VALUES (?,?,?,?,?,?)`,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, userID := range userIDs {
		id, err := r.generateID()
		if err != nil {
			return nil, err
		}

		if _, err := stmt.Exec(id, userID, studyGroupID, authority, now, now); err != nil {
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}
