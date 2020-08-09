package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateDepartment(organizationID uint64, name string, url string) (*model.Department, error) {
	return app.repo.CreateDepartment(organizationID, name, url)
}
