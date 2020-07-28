package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateMembership(studyGroupID uint64, userName string) (*model.Membership, error) {
	user, err := app.repo.FindUserByUserName(userName)
	if err != nil {
		return nil, err
	}

	return app.repo.CreateMembership(user.ID, studyGroupID, 2)
}
