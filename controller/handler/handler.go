package handler

import (
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/config"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/controller"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/database"

	m "github.com/PKL-Angkasa-Pura-I/backend-pkl/middleware"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/repository"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewService(repo, conf)

	cont := controller.EchoController{
		Svc: svc,
	}

	e.GET("/pkl_v1/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	e.POST("/pkl_v1/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	api := e.Group("/pkl_v1", middleware.CORS())

	m.LogMiddleware(e)
	api.POST("/admins/login", cont.LoginAdminController)

	api.POST("/divisions", cont.CreateDivisionController)
}