package service

import (
	"errors"

	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateStudyGroup(name string, url string) (*model.StudyGroup, error) {
	if name == "" {
		return nil, errors.New("empty study group name")
	}
	if url == "" {
		return nil, errors.New("empty url")
	}

	return app.repo.CreateStudyGroup(name, url)
}
