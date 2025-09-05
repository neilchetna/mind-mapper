package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/internal/repository"
	"gorm.io/gorm"
)

type NodeRepository interface {
	GetNodes(ctx context.Context, chartId uuid.UUID, userId uuid.UUID) ([]models.Node, error)
	CreatNode(ctx context.Context, node *models.Node) error
}

type NodeService struct {
	node NodeRepository
	edge EdgeRepository
}

func (svc *NodeService) GetNodes(ctx context.Context, chartID uuid.UUID, userID uuid.UUID) ([]models.Node, error) {
	return svc.node.GetNodes(ctx, chartID, userID)
}

func (svc *NodeService) CreateNode(ctx context.Context, node *models.Node) (models.Edge, error) {
	var edge models.Edge
	err := svc.node.CreatNode(ctx, node)
	if err != nil {
		return edge, err
	}
	edge.ChartId = node.ChartId
	edge.Source = node.ParentId
	edge.Target = node.ID

	err = svc.edge.CreateEdge(ctx, &edge)
	if err != nil {
		return edge, err
	}
	return edge, nil
}

func NodeServiceBuilder(db *gorm.DB) *NodeService {
	nodeRepo := repository.NodeRepositoryBuilder(db)
	edgeRepo := repository.EdgeRepositoryBuilder(db)
	return &NodeService{node: nodeRepo, edge: edgeRepo}
}
