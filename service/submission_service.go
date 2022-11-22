package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	nanoid "github.com/aidarkhanov/nanoid/v2"
)

func (s *svc) CreateSubmissionService(submission model.Submission) (string, error) {
	if submission.Name == "" || submission.Email == "" || submission.SchoolOrigin == "" {
		return "", fmt.Errorf("insert submission error")
	}

	code, _ := nanoid.New()
	submission.CodeSubmission = code

	err := s.repo.CreateSubmission(submission)
	if err != nil {
		return "", err
	}

	res, err := s.repo.GetSubmissionByCodeSubmission(submission.CodeSubmission)
	if err != nil {
		return "", err
	}

	dt := res.CreatedAt.Format("01-02")
	format_dt := strings.ReplaceAll(dt, "-", "")
	int_dt, _ := strconv.Atoi(format_dt)
	res.CodeSubmission = fmt.Sprintf("%s%d", "P-", int_dt+int(res.DivisionID)+int(res.Study_fieldID)+int(res.ID))
	err = s.repo.UpdateSubmissionByID(int(res.ID), res)
	if err != nil {
		return "", err
	}

	return res.CodeSubmission, nil
}
