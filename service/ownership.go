package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateOwnership(studyGroupID uint64, eventID uint64) (*model.Ownership, error) {
	return app.repo.CreateOwnership(studyGroupID, eventID)
}
