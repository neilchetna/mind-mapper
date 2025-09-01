package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/models"
	"gorm.io/gorm"
)

type ChartRepository struct {
	db *gorm.DB
}

func ChartRepositoryBuilder(db *gorm.DB) *ChartRepository {
	return &ChartRepository{db}
}

func (r *ChartRepository) CreateChart(ctx context.Context, chart *models.Chart) error {
	res := r.db.WithContext(ctx).Create(chart)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *ChartRepository) GetCharts(ctx context.Context, userId uuid.UUID) ([]models.Chart, error) {
	var charts []models.Chart
	res := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&charts)

	if res.Error != nil {
		return nil, res.Error
	}

	return charts, nil
}

func (r *ChartRepository) GetChartById(ctx context.Context, chartId uuid.UUID, userId uuid.UUID) (models.Chart, error) {
	var chart models.Chart
	res := r.db.WithContext(ctx).Where("user_id = ?", userId).First(&chart, "id = ?", chartId)

	if res.Error != nil {
		return chart, res.Error
	}

	return chart, nil
}
