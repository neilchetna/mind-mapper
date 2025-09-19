package rest

import (
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/neilchetna/mind-mapper/internal/rest/handlers"
	"github.com/neilchetna/mind-mapper/internal/rest/middleware"
	"github.com/neilchetna/mind-mapper/internal/service"
	"gorm.io/gorm"
)

func BuildRoutes(e *echo.Echo, db *gorm.DB, q *asynq.Client) {
	bindMiddleware(e, db)

	e.GET("/", handlers.RootHanlder)

	// Charts api
	chartGroup := e.Group("/charts")
	chartSvc := service.ChartServiceBuilder(db)
	handlers.NewChartHanlder(chartGroup, chartSvc)

	// Nodes group
	nodeGroup := e.Group("/charts/:chartId/nodes")
	nodeSvc := service.NodeServiceBuilder(db, q)
	handlers.NewNodeHandlerBuilder(nodeGroup, nodeSvc)

	// Edges group
	edgeGroup := e.Group("/charts/:chartId/edges")
	edgeSVC := service.NewEdgeServiceBuilder(db)
	handlers.NewEdgeHandlerBuilder(edgeGroup, edgeSVC)

}

func bindMiddleware(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.CORS())
	e.Use(middleware.Authenticate)
	e.Use(middleware.SyncUser(db))
}
