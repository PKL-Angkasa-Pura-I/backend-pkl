package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) CreatePivotDivisionField(pivot_division_field model.Pivot_division_field) error {
	res := r.DB.Create(&pivot_division_field)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert new division field")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllDivisionField(division_id int) []model.List_division_field {
	res := []model.List_division_field{}
	r.DB.Model(&model.Pivot_division_field{}).Select("study_fields.name").
		Joins("JOIN divisions on divisions.id = pivot_division_fields.division_id").
		Joins("JOIN study_fields on study_fields.id = pivot_division_fields.study_field_id").
		Where("pivot_division_fields.division_id = ?", division_id).
		Scan(&res)

	return res
}

func (r *repositoryMysqlLayer) DeleteOnePivotDivisionField(division_id, study_field_id int) error {
	res := r.DB.Unscoped().Where("division_id = ? AND study_field_id = ?", division_id, study_field_id).Delete(&model.Pivot_division_field{})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete, division and study field not found")
	}

	return nil
}

func (r *repositoryMysqlLayer) CheckDivisionByID(id int) (pivot_division_field model.Pivot_division_field, err error) {
	res := r.DB.Where("division_id = ?", id).Find(&pivot_division_field)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("division not found")
	}

	return
}

func (r *repositoryMysqlLayer) CheckStudyFieldByID(id int) (pivot_division_field model.Pivot_division_field, err error) {
	res := r.DB.Where("study_field_id = ?", id).Find(&pivot_division_field)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("study field not found")
	}

	return
}
