package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/internal/repository"
	"gorm.io/gorm"
)

type NodeRepository interface {
	GetNode(ctx context.Context, chartId uuid.UUID, userId uuid.UUID) ([]models.Node, error)
	CreatNode(ctx context.Context, node *models.Node) error
}

type NodeService struct {
	r NodeRepository
}

func (svc *NodeService) GetNodes(ctx context.Context, chartID uuid.UUID, userID uuid.UUID) ([]models.Node, error) {
	return svc.r.GetNode(ctx, chartID, userID)
}

func (svc *NodeService) CreateNode(ctx context.Context, node *models.Node) error {
	return svc.r.CreatNode(ctx, node)
}

func NodeServiceBuilder(db *gorm.DB) *NodeService {
	nodeRepo := repository.NodeRepositoryBuilder(db)
	return &NodeService{r: nodeRepo}
}
