package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) CreateStudyField(study_field model.Study_field) error {
	res := r.DB.Create(&study_field)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert study field")
	}

	return nil
}
