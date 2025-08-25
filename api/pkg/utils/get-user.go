package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/neilchetna/mind-mapper/internal/models"
)

func GetReqUser(c echo.Context) *models.User {	
	return c.Get("user").(*models.User)
}