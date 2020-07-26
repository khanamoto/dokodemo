package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var studyGroupNotFoundError = model.NotFoundError("study_group")

func (r *repository) CreateStudyGroup(name string, url string) (*model.StudyGroup, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO study_group (id, name, url, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		id, name, url, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.StudyGroup{ID: id, Name: name, URL: url}, nil
}

// func (r *repository) FindStudyGroupByURL(url string) (*model.StudyGroup, error) {
// 	var studyGroup model.StudyGroup
// 	err := r.db.Get(
// 		&studyGroup,
// 		`SELECT id,name,url FROM study_group WHERE url = ? LIMIT 1`, url,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, studyGroupNotFoundError
// 		}
// 		return nil, err
// 	}
// 	return &studyGroup, nil
// }
