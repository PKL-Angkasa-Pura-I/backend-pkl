package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) CreateDivision(division model.Division) error {
	res := r.DB.Create(&division)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert division")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllDivision() []model.Division {
	divisions := []model.Division{}
	r.DB.Find(&divisions)

	return divisions
}
