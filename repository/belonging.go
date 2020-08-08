package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var belongingNotFoundError = model.NotFoundError("belonging")

func (r *repository) CreateBelonging(userID uint64, organizationID uint64, authority int32) (*model.Belonging, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO belonging (id, user_id, organization_id, authority, created_at, updated_at) VALUES (?,?,?,?,?,?)`,
		id, userID, organizationID, authority, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Belonging{ID: id, UserID: userID, OrganizationID: organizationID, Authority: authority}, nil
}
