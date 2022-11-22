package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"gorm.io/gorm/clause"
)

func (r *repositoryMysqlLayer) CreateSubmission(submission model.Submission) error {
	res := r.DB.Create(&submission)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert submission")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetSubmissionByCodeSubmission(code_submission string) (submission model.Submission, err error) {
	res := r.DB.Where("code_submission = ?", code_submission).Preload(clause.Associations).Find(&submission)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("code submission not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateSubmissionByID(id int, submission model.Submission) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&submission)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update submission")
	}

	return nil
}
