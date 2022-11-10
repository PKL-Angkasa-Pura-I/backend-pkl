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

	LoginAdmin(username, password string) (string, int)
	GetAdminByUsernameService(username string) (model.Admin, error)

	ClaimToken(bearer *jwt.Token) string
}
