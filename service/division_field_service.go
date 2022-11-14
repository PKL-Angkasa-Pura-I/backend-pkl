package service

import (
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (s *svc) CreatePivotDivisionFieldService(pivot_division_field model.Pivot_division_field) error {
	return s.repo.CreatePivotDivisionField(pivot_division_field)
}

func (s *svc) GetAllDivisionFieldService(division_id int) []model.List_division_field {
	return s.repo.GetAllDivisionField(division_id)
}

func (s *svc) DeleteOnePivotDivisionFieldService(division_id, study_field_id int) error {
	return s.repo.DeleteOnePivotDivisionField(division_id, study_field_id)
}
