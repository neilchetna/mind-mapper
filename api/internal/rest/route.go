package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/neilchetna/mind-mapper/internal/rest/handlers"
	"github.com/neilchetna/mind-mapper/internal/rest/middleware"
	"github.com/neilchetna/mind-mapper/internal/service"
	"gorm.io/gorm"
)

func BuildRoutes(e *echo.Echo, db *gorm.DB) {
	bindMiddleware(e, db)

	e.GET("/", handlers.RootHanlder)

	// Charts api
	chartGroup := e.Group("/charts")
	chartSvc := service.ChartServiceBuilder(db)
	handlers.NewChartHanlder(chartGroup, chartSvc)

	// Nodes group
	nodeGroup := e.Group("/charts/:chartId/nodes")
	nodeSvc := service.NodeServiceBuilder(db)
	handlers.NewNodeHandlerBuilder(nodeGroup, nodeSvc)
}

func bindMiddleware(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.CORS())
	e.Use(middleware.Authenticate)
	e.Use(middleware.SyncUser(db))
}
