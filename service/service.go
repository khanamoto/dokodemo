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
	FindUserByUserName(userName string) (*model.User, error)
	LoginUser(userName string, password string) (bool, error)
	CreateNewToken(userID uint64, expiresAt time.Time) (string, error)

	CreateStudyGroup(name string, url string) (*model.StudyGroup, error)

	CreateSubStudyGroup(studyGroupID uint64, name string, url string) (*model.SubStudyGroup, error)

	CreateMembership(studyGroupID uint64, userName string) (*model.Membership, error)

	CreateSubMembership(subStudyGroupID uint64, userName []string) (*model.SubMembership, error)
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
