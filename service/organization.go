package service

import (
	"github.com/khanamoto/dokodemo/model"
)

func (app *dokodemo) CreateOrganization(name string, url string) (*model.Organization, error) {
	return app.repo.CreateOrganization(name, url)
}
