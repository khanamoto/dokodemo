package service

import (
	"errors"

	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateMembership(name string, url string, userName string) (*model.Membership, error) {
	// nameとurlのチェック
	if name == "" {
		return nil, errors.New("empty study group name")
	}
	if url == "" {
		return nil, errors.New("empty url")
	}
	// 一人のuserを検索
	user, err := app.repo.FindUserByUserName(userName)
	if err != nil {
		return nil, err
	}
	// studyGroupを作成
	studyGroup, err := app.repo.CreateStudyGroup(name, url)
	if err != nil {
		return nil, err
	}
	// TODO: 更新で使う
	// if err := app.repo.FindMembership(user.ID, studyGroup.ID); err != nil {
	// 	return err
	// }
	return app.repo.CreateMembership(user.ID, studyGroup.ID, 2)
}
