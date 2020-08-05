package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var subStudyGroupNotFoundError = model.NotFoundError("sub_study_group")

func (r *repository) CreateSubStudyGroup(studyGroupID uint64, name string, url string) (*model.SubStudyGroup, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO sub_study_group (id, study_group_id, name, url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`,
		id, studyGroupID, name, url, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.SubStudyGroup{ID: id, StudyGroupID: studyGroupID, Name: name, URL: url}, nil
}
