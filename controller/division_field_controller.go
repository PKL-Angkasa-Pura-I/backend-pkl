package controller

import (
	"strconv"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreatePivotDivisionFieldController(c echo.Context) error {

	id := c.Param("id_division")
	id_int, _ := strconv.Atoi(id)

	pivot_division_field := model.Pivot_division_field{}
	if err := c.Bind(&pivot_division_field); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	res, err := ce.Svc.GetDivisionByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "Division not found",
		})
	}

	res2, err := ce.Svc.GetStudyFieldByIDService(int(pivot_division_field.Study_fieldID))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "Study Field not found",
		})
	}

	pivot_division_field.DivisionID = uint(id_int)
	pivot_division_field.Division = res
	pivot_division_field.Study_field = res2

	err = ce.Svc.CreatePivotDivisionFieldService(pivot_division_field)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":         "success",
		"division_name":    pivot_division_field.Division.Name,
		"study_field_name": pivot_division_field.Study_field.Name,
	})
}
