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
	CreateNewToken(userID uint64, token string, expiresAt time.Time) error

	CreateStudyGroup(name string, url string) (*model.StudyGroup, error)
	// FindStudyGroupByURL(url string) (*model.StudyGroup, error)

	CreateMembership(userID uint64, studyGroupID uint64, authority int32) (*model.Membership, error)

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
