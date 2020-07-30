package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var userNotFoundError = model.NotFoundError("user")

func (r *repository) CreateNewUser(name string, userName string, email string, passwordHash string) error {
	id, err := r.generateID()
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO user (id, name, user_name, email, password_hash, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		id, name, userName, email, passwordHash, now, now,
	)
	return err
}

func (r *repository) FindUserByUserName(userName string) (*model.User, error) {
	var user model.User
	err := r.db.Get(
		&user,
		`SELECT id, user_name FROM user WHERE user_name = ? LIMIT 1`, userName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userNotFoundError
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindPasswordHashByUserName(userName string) (string, error) {
	var hash string
	err := r.db.Get(
		&hash,
		`SELECT password_hash FROM user WHERE user_name = ? LIMIT 1`, userName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", nil
	}
	return hash, nil
}

func (r *repository) CreateNewToken(userID uint64, token string, expiresAt time.Time) error {
	now := time.Now()
	_, err := r.db.Exec(
		`INSERT INTO user_session (user_id, token, expires_at, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		userID, token, expiresAt, now, now,
	)
	return err
}
