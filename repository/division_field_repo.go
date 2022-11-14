package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) CreatePivotDivisionField(pivot_division_field model.Pivot_division_field) error {
	res := r.DB.Create(&pivot_division_field)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert new division field")
	}

	return nil
}
