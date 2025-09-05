package handlers

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/pkg/utils"
)

type EdgeService interface {
	GetEdges(ctx context.Context, chartID uuid.UUID) ([]models.Edge, error)
}

type EdgeHandler struct {
	svc EdgeService
}

func NewEdgeHandlerBuilder(g *echo.Group, svc EdgeService) {
	handler := &EdgeHandler{svc}

	g.GET("", handler.Query)
}

func (h *EdgeHandler) Query(c echo.Context) error {
	chartId, err := utils.ParseUUIDParam(c, "chartId")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	edges, err := h.svc.GetEdges(ctx, chartId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, edges)
}
