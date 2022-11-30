package service

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (s *svc) CreateDivisionService(division model.Division) error {
	if division.Name == "" {
		return fmt.Errorf("error insert division")
	}
	return s.repo.CreateDivision(division)
}

func (s *svc) GetAllDivisionService() []model.Division {
	return s.repo.GetAllDivision()
}

func (s *svc) GetDivisionByIDService(id int) (model.Division, error) {
	return s.repo.GetDivisionByID(id)
}

func (s *svc) UpdateDivisionByIDService(id int, division model.Division) error {
	return s.repo.UpdateDivisionByID(id, division)
}

func (s *svc) DeleteDivisionByIDService(id int) error {
	s.repo.DeleteAllDivisionField(id)

	return s.repo.DeleteDivisionByID(id)
}
