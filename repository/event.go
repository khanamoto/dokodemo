package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var eventNotFoundError = model.NotFoundError("event")

func (r *repository) CreateEvent(name string, eventDate time.Time, description string, place string) (*model.Event, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO event (id, name, event_date, description, place, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		id, name, eventDate, description, place, now, now,
	)
	return &model.Event{ID: id, Name: name, EventDate: eventDate, Description: description, Place: place}, nil
}
