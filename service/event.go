package service

import (
	"time"

	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateEvent(name string, eventDate time.Time, description string, place string) (*model.Event, error) {
	return app.repo.CreateEvent(name, eventDate, description, place)
}
