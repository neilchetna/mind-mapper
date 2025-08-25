package handlers

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/pkg/utils"
)

type ChartService interface {
	CreateChart(ctx context.Context, chart *models.Chart) error
	GetCharts(ctx context.Context, userId uuid.UUID) ([]models.Chart, error)
}

type ChartHandler struct {
	Service ChartService
} 

func NewChartHanlder(g *echo.Group, svc ChartService) {
	handler := &ChartHandler{Service: svc}

	g.POST("", handler.Create)
	g.GET("", handler.Query)
}

func (h *ChartHandler) Create(c echo.Context) error {
	var chart models.Chart

	err := c.Bind(&chart)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	user := utils.GetReqUser(c) 
	chart.UserId = user.ID

	ctx := c.Request().Context()
	err = h.Service.CreateChart(ctx, &chart)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, chart)
}

func (h *ChartHandler) Query(c echo.Context) error {
	ctx := c.Request().Context()
	var charts []models.Chart

	user := utils.GetReqUser(c)
	
	charts, err := h.Service.GetCharts(ctx, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}


	return c.JSON(http.StatusOK, charts)
}