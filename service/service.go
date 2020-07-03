package service

import (
	"math/rand"
	"time"

	"github.com/khanamoto/dokodemo/model"
	"github.com/khanamoto/dokodemo/repository"
	"github.com/khanamoto/dokodemo/titleFetcher"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Dokodemo interface {
	 Close() error

	 CreateNewUser(name string, passwordHash string) error
	 FindUserByName(name string) (*model.User, error)
	 FindUserByID(userID uint64) (*model.User, error)
	 ListUsersByIDs(userIDs []uint64) ([]*model.User, error)
	 LoginUser(name string, password string) (bool, error)
	 CreateNewToken(userID uint64, expiresAt time.Time) (string, error)
	 FindUserByToken(token string) (*model.User, error)
}

func NewApp(repo repository.Repository, titleFetcher titleFetcher.TitleFetcher) Dokodemo {
	return &dokodemo{repo, titleFetcher }
}

type dokodemo struct {
	repo         repository.Repository
	titleFetcher titleFetcher.TitleFetcher
}

func (app *dokodemo) Close() error {
	return app.repo.Close()
}