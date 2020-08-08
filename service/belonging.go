package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateBelonging(organizationID uint64, userName string) (*model.Belonging, error) {
	user, err := app.repo.FindUserByUserName(userName)
	if err != nil {
		return nil, err
	}

	return app.repo.CreateBelonging(user.ID, organizationID, 2)
}
