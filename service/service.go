package service

import (
	"math/rand"
	"time"

	"github.com/khanamoto/dokodemo/model"
	"github.com/khanamoto/dokodemo/repository"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Dokodemo interface {
	Close() error

	CreateNewUser(name string, userName string, email string, passwordHash string) error
	FindUserByName(name string) (*model.User, error)
	// FindUserByID(userID uint64) (*model.User, error)
	// ListUsersByIDs(userIDs []uint64) ([]*model.User, error)
	// LoginUser(name string, password string) (bool, error)
	CreateNewToken(userID uint64, expiresAt time.Time) (string, error)
	// FindUserByToken(token string) (*model.User, error)
}

func NewApp(repo repository.Repository) Dokodemo {
	return &dokodemo{repo}
}

type dokodemo struct {
	repo repository.Repository
}

func (app *dokodemo) Close() error {
	return app.repo.Close()
}
