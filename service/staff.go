package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateStaff(departmentID uint64, userNames []string) (*model.Staff, error) {
	cap := len(userNames)
	userIDs := make([]uint64, cap, cap)
	for i := 0; i < cap; i++ {
		user, err := app.FindUserByUserName(userNames[i])
		if err != nil {
			return nil, err
		}
		userIDs[i] = user.ID
	}

	return app.repo.CreateStaff(userIDs, departmentID, 2)
}
