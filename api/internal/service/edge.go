package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/internal/repository"
	"gorm.io/gorm"
)

type EdgeRepository interface {
	CreateEdge(ctx context.Context, edge *models.Edge) error
	GetEdges(ctx context.Context, chartId uuid.UUID) ([]models.Edge, error)
}

type EdgeService struct {
	r EdgeRepository
}

func (svc *EdgeService) GetEdges(ctx context.Context, chartId uuid.UUID) ([]models.Edge, error) {
	return svc.r.GetEdges(ctx, chartId)
}

func NewEdgeServiceBuilder(db *gorm.DB) *EdgeService {
	r := repository.EdgeRepositoryBuilder(db)
	return &EdgeService{r}
}
