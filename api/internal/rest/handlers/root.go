package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RootHanlder(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Mind mapper api is working",
		"version": "0.1.0",
		"status":  "running",
	})
}
