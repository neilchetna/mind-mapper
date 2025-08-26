package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/internal/repository"
	"gorm.io/gorm"
)

type ChartRepository interface {
	CreateChart(ctx context.Context, chart *models.Chart) error
	GetCharts(ctx context.Context, userId uuid.UUID) ([]models.Chart, error)
	GetChartById(ctx context.Context, chartId uuid.UUID, userId uuid.UUID) (models.Chart, error)
} 

type ChartService struct {
	r ChartRepository
}

func (svc *ChartService) CreateChart(ctx context.Context, chart *models.Chart) error {
	return svc.r.CreateChart(ctx, chart)
}

func (svc *ChartService) GetCharts(ctx context.Context, userId uuid.UUID) ([]models.Chart, error) {
	return svc.r.GetCharts(ctx, userId)
}

func (svc *ChartService) GetChart(ctx context.Context, chartId uuid.UUID, userId uuid.UUID) (models.Chart, error) {
	return svc.r.GetChartById(ctx, chartId, userId)
} 

func ChartServiceBuilder(db *gorm.DB) *ChartService {
	chartRepo := repository.ChartRepositoryBuilder(db)
	return &ChartService{r: chartRepo}
}