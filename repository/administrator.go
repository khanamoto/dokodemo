package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var administratorNotFoundError = model.NotFoundError("administrator")

func (r *repository) CreateAdministrator(userID uint64, eventID uint64) (*model.Administrator, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO administrator (id, user_id, event_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		id, userID, eventID, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Administrator{ID: id, UserID: userID, EventID: eventID}, nil
}
