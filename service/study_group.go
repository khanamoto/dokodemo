package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateStudyGroup(departmentID uint64, name string, url string) (*model.StudyGroup, error) {
	return app.repo.CreateStudyGroup(departmentID, name, url)
}
