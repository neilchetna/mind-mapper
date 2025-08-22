package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/neilchetna/mind-mapper/internal/rest/handlers"
	"gorm.io/gorm"
)

func BuildRoutes(e *echo.Echo, db *gorm.DB) {

	e.GET("/", handlers.RootHanlder)
}
