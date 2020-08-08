package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var organizationNotFoundError = model.NotFoundError("organization")

func (r *repository) CreateOrganization(name string, url string) (*model.Organization, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO organization (id, name, url, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		id, name, url, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Organization{ID: id, Name: name, URL: url}, nil
}
