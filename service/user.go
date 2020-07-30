package service

import (
	"errors"
	"math/rand"
	"time"

	"github.com/khanamoto/dokodemo/model"
	"golang.org/x/crypto/bcrypt"
)

func (app *dokodemo) CreateNewUser(name string, userName string, email string, password string) (err error) {
	if name == "" {
		return errors.New("empty user name")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return app.repo.CreateNewUser(name, userName, email, string(passwordHash))
}

func (app *dokodemo) LoginUser(userName string, password string) (bool, error) {
	passwordHash, err := app.repo.FindPasswordHashByUserName(userName)
	if err != nil {
		return false, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (app *dokodemo) FindUserByUserName(userName string) (*model.User, error) {
	return app.repo.FindUserByUserName(userName)
}

func generateToken() string {
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_@"
	l := len(table)
	ret := make([]byte, 128)
	src := make([]byte, 128)
	rand.Read(src)
	for i := 0; i < 128; i++ {
		ret[i] = table[int(src[i])%l]
	}
	return string(ret)
}

func (app *dokodemo) CreateNewToken(userID uint64, expiresAt time.Time) (string, error) {
	token := generateToken()
	err := app.repo.CreateNewToken(userID, token, expiresAt)
	if err != nil {
		return "", err
	}
	return token, nil
}
