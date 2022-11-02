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