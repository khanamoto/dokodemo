package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateMembership(studyGroupID uint64, userNames []string) (*model.Membership, error) {
	cap := len(userNames)
	userIDs := make([]uint64, cap, cap)
	for i := 0; i < cap; i++ {
		user, err := app.FindUserByUserName(userNames[i])
		if err != nil {
			return nil, err
		}
		userIDs[i] = user.ID
	}

	return app.repo.CreateMembership(userIDs, studyGroupID, 2)
}
