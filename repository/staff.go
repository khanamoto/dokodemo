package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var staffNotFoundError = model.NotFoundError("staff")

func (r *repository) CreateStaff(userIDs []uint64, departmentID uint64, authority int32) (*model.Staff, error) {
	now := time.Now()

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		`INSERT INTO staff (id, user_id, department_id, authority, created_at, updated_at) VALUES (?,?,?,?,?,?)`,
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

		if _, err := stmt.Exec(id, userID, departmentID, authority, now, now); err != nil {
			return nil, err
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return nil, nil
}
