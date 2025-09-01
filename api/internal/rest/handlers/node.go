package handlers

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/pkg/utils"
)

type NodeService interface {
	GetNodes(ctx context.Context, chartID uuid.UUID, userID uuid.UUID) ([]models.Node, error)
}

type NodeHandler struct {
	svc NodeService
}

func NewNodeHandlerBuilder(g *echo.Group, svc NodeService) {
	handler := &NodeHandler{svc}

	g.GET("", handler.Query)
}

func (h *NodeHandler) Query(c echo.Context) error {
	chartID, err := utils.ParseUUIDParam(c, "chartId")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	user := utils.GetReqUser(c)
	nodes, err := h.svc.GetNodes(ctx, chartID, user.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nodes)
}
