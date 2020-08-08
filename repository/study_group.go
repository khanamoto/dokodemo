package repository

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/khanamoto/dokodemo/model"
)

var studyGroupNotFoundError = model.NotFoundError("study_group")

func (r *repository) CreateStudyGroup(departmentID uint64, name string, url string) (*model.StudyGroup, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO study_group (id, department_id, name, url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`,
		id, departmentID, name, url, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.StudyGroup{ID: id, DepartmentID: departmentID, Name: name, URL: url}, nil
}
