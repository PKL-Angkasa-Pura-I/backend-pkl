package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) CreateTrainee(trainee model.Trainee) error {
	res := r.DB.Create(&trainee)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert trainee repo")
	}

	return nil
}

func (r *repositoryMysqlLayer) CountTrainee(submission_id int) int {
	var count int64
	r.DB.Model(&model.Trainee{}).Where("submission_id = ?", submission_id).Count(&count)

	return int(count)
}
