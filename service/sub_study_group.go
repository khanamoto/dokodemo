package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateSubStudyGroup(studyGroupID uint64, name string, url string) (*model.SubStudyGroup, error) {
	return app.repo.CreateSubStudyGroup(studyGroupID, name, url)
}
