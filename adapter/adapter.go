package adapter

import "github.com/PKL-Angkasa-Pura-I/backend-pkl/model"

type AdapterRepository interface {
	CreateDivision(division model.Division) error

	GetAdminByUsername(username string) (admin model.Admin, err error)
}

type AdapterService interface {
	CreateDivisionService(division model.Division) error

	LoginAdmin(username, password string) (string, int)
}
