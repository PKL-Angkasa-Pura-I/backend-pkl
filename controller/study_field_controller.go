package controller

import (
	"strconv"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateStudyFieldController(c echo.Context) error {

	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	study_field := model.Study_field{}
	if err := c.Bind(&study_field); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.CreateStudyFieldService(study_field)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":    "success",
		"study_field": study_field.Name,
	})
}

func (ce *EchoController) GetAllStudyFieldController(c echo.Context) error {

	study_fields := ce.Svc.GetAllStudyFieldService()

	return c.JSON(200, map[string]interface{}{
		"messages":    "success",
		"study_field": study_fields,
	})
}

func (ce *EchoController) GetOneStudyFieldController(c echo.Context) error {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	res, err := ce.Svc.GetStudyFieldByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "study field not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages":    "success",
		"study_field": res,
	})
}
