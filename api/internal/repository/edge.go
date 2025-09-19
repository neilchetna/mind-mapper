package repository

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/models"
	"gorm.io/gorm"
)

type EdgeRepository struct {
	db *gorm.DB
}

func EdgeRepositoryBuilder(db *gorm.DB) *EdgeRepository {
	return &EdgeRepository{db}
}

func (r *EdgeRepository) CreateEdge(ctx context.Context, edge *models.Edge) error {
	res := r.db.WithContext(ctx).Create(edge)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *EdgeRepository) GetEdges(ctx context.Context, chartId uuid.UUID) ([]models.Edge, error) {
	var edges []models.Edge
	res := r.db.WithContext(ctx).Where(models.Edge{ChartId: chartId}).Find(&edges)

	if res.Error != nil {
		return nil, res.Error
	}

	return edges, nil
}

func (r *EdgeRepository) GetEdgeByData(ctx context.Context, edge *models.Edge) error {
	log.Print(edge)
	res := r.db.WithContext(ctx).Where(edge).First(&models.Edge{})

	if res.Error != nil {
		return res.Error
	}

	return nil
}
