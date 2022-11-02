package controller

import (
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateDivisionController(c echo.Context) error {

	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	}

	division := model.Division{}
	if err := c.Bind(&division); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.CreateDivisionService(division)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":      "success",
		"division_name": division.Name,
	})
}

func (ce *EchoController) GetAllDivisionController(c echo.Context) error {

	divisions := ce.Svc.GetAllDivisionService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"division": divisions,
	})
}
