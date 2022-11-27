package service

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (s *svc) CreateTraineeService(trainee model.Trainee) error {
	if trainee.Name == "" || trainee.Email == "" || trainee.Trainee_Student_id == "" || trainee.Phone == "" || trainee.Jurusan == "" || trainee.Gender == "" {
		return fmt.Errorf("error insert trainee service")
	}
	return s.repo.CreateTrainee(trainee)
}

func (s *svc) CountTraineeService(submission_id int) int {
	return s.repo.CountTrainee(submission_id)
}
