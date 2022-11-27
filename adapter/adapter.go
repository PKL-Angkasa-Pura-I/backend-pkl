package adapter

import (
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"github.com/golang-jwt/jwt"
)

type AdapterRepository interface {
	CreateDivision(division model.Division) error
	GetAllDivision() []model.Division
	GetDivisionByID(id int) (division model.Division, err error)
	UpdateDivisionByID(id int, division model.Division) error
	DeleteDivisionByID(id int) error

	CreateStudyField(study_field model.Study_field) error
	GetAllStudyField() []model.Study_field
	GetStudyFieldByID(id int) (study_field model.Study_field, err error)
	UpdateStudyFieldByID(id int, study_field model.Study_field) error
	DeleteStudyFieldByID(id int) error

	CreatePivotDivisionField(pivot_division_field model.Pivot_division_field) error
	GetAllDivisionField(division_id int) []model.List_division_field
	DeleteOnePivotDivisionField(division_id, study_field_id int) error
	CheckPivotDivisionByID(id int) (pivot_division_field model.Pivot_division_field, err error)
	CheckPivotStudyFieldByID(id int) (pivot_division_field model.Pivot_division_field, err error)
	GetAllDivisionStudyField() []model.Pivot_division_field
	GetDivisionOnPivot(id int) []model.Pivot_division_field
	CheckDivisonField(id_division, id_study_field int) (pivot_division_field model.Pivot_division_field, err error)

	CreateSubmission(submission model.Submission) error
	GetSubmissionByCodeSubmission(code_submission string) (submission model.Submission, err error)
	GetSubmissionByID(id int) (submission model.Submission, err error)
	UpdateSubmissionByID(id int, submission model.Submission) error
	GetAllSubmission() []model.Submission

	GetAdminByUsername(username string) (admin model.Admin, err error)
}

type AdapterService interface {
	CreateDivisionService(division model.Division) error
	GetAllDivisionService() []model.Division
	GetDivisionByIDService(id int) (model.Division, error)
	UpdateDivisionByIDService(id int, division model.Division) error
	DeleteDivisionByIDService(id int) error

	CreateStudyFieldService(study_field model.Study_field) error
	GetAllStudyFieldService() []model.Study_field
	GetStudyFieldByIDService(id int) (model.Study_field, error)
	UpdateStudyFieldByIDService(id int, study_field model.Study_field) error
	DeleteStudyFieldByIDService(id int) error

	CreatePivotDivisionFieldService(pivot_division_field model.Pivot_division_field) error
	GetAllDivisionFieldService(division_id int) []model.List_division_field
	DeleteOnePivotDivisionFieldService(division_id, study_field_id int) error
	GetAllDivisionStudyFieldService() []model.List_pivot

	CreateSubmissionService(submission model.Submission) (string, error)
	GetAllSubmissionService() []model.Submission
	GetSubmissionByCodeSubmissionService(code_submission string) (model.Submission, error)
	GetSubmissionByIDService(id int) (model.Submission, error)
	UpdateSubmissionByIDService(id int, submission model.Submission) error

	LoginAdmin(username, password string) (string, int)
	GetAdminByUsernameService(username string) (model.Admin, error)

	ClaimToken(bearer *jwt.Token) string
}
