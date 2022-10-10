package adapter

import "github.com/PKL-Angkasa-Pura-I/backend-pkl/model"

type AdapterRepository interface {
	CreateDivision(division model.Division) error
}

type AdapterService interface {
	CreateDivisionService(division model.Division) error
}
