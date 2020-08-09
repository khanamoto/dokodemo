package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var departmentNotFoundError = model.NotFoundError("department")

func (r *repository) CreateDepartment(organizationID uint64, name string, url string) (*model.Department, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO department (id, organization_id, name, url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`,
		id, organizationID, name, url, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Department{ID: id, OrganizationID: organizationID, Name: name, URL: url}, nil
}
