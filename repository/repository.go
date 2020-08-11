package repository

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/khanamoto/dokodemo/model"
)

type Repository interface {
	CreateNewUser(name string, userName string, email string, passwordHash string) error
	FindUserByUserName(userName string) (*model.User, error)
	FindPasswordHashByUserName(userName string) (string, error)
	CreateNewToken(userID uint64, token string, expiresAt time.Time) error
	FindUserByToken(token string) (*model.User, error)

	CreateOrganization(name string, url string) (*model.Organization, error)
	CreateBelonging(userID uint64, organizationID uint64, authority int32) (*model.Belonging, error)

	CreateDepartment(organizationID uint64, name string, url string) (*model.Department, error)
	CreateStaff(userIDs []uint64, departmentID uint64, authority int32) (*model.Staff, error)

	CreateStudyGroup(departmentID uint64, name string, url string) (*model.StudyGroup, error)
	CreateMembership(userIDs []uint64, studyGroupID uint64, authority int32) (*model.Membership, error)

	CreateEvent(name string, eventDate time.Time, description string, place string) (*model.Event, error)
	CreateOwnership(studyGroupID uint64, eventID uint64) (*model.Ownership, error)
	CreateAdministrator(userID uint64, eventID uint64) (*model.Administrator, error)

	Close() error
}

func New(dsn string) (Repository, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Opening mysql failed: %v", err)
	}
	return &repository{db: db}, nil
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) generateID() (uint64, error) {
	var id uint64
	err := r.db.Get(&id, "SELECT UUID_SHORT()")
	return id, err
}

func (r *repository) Close() error {
	return r.db.Close()
}
