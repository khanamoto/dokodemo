package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateAdministrator(userID uint64, eventID uint64) (*model.Administrator, error) {
	return app.repo.CreateAdministrator(userID, eventID)
}
