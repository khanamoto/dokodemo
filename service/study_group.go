package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateStudyGroup(name string, url string) (*model.StudyGroup, error) {
	// TODO: 更新時使用？
	// if _, err := app.repo.FindStudyGroupByURL(url); err != nil {
	// 	return nil, err
	// }

	return app.repo.CreateStudyGroup(name, url)
}
