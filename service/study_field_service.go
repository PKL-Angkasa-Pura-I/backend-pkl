package service

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (s *svc) CreateStudyFieldService(study_field model.Study_field) error {
	if study_field.Name == "" {
		return fmt.Errorf("error insert study field")
	}
	return s.repo.CreateStudyField(study_field)
}
