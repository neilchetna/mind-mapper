package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/internal/repository"
	"gorm.io/gorm"
)

type NodeRepository interface {
	GetUserNodes(ctx context.Context, chartId uuid.UUID, userId uuid.UUID) ([]models.Node, error)
	GetNodes(ctx context.Context, chartId uuid.UUID) ([]models.Node, error)
	CreateNode(ctx context.Context, node *models.Node) error
	CreateBulkNodes(ctx context.Context, nodes *[]models.Node) error
}

type NodeService struct {
	node NodeRepository
	edge EdgeRepository
	q    *QueueService
}

func (svc *NodeService) GetNodes(ctx context.Context, chartID uuid.UUID, userID uuid.UUID) ([]models.Node, error) {
	return svc.node.GetUserNodes(ctx, chartID, userID)
}

func (svc *NodeService) GetChartNodes(ctx context.Context, chartID uuid.UUID) ([]models.Node, error) {
	return svc.node.GetNodes(ctx, chartID)
}

func (svc *NodeService) CreateNode(ctx context.Context, node *models.Node) (models.Edge, error) {
	var edge models.Edge
	err := svc.node.CreateNode(ctx, node)
	if err != nil {
		return edge, err
	}

	edge = models.Edge{Target: node.ID, Source: node.ParentId, ChartId: node.ChartId}
	if err := svc.edge.GetEdgeByData(ctx, &edge); err != nil {
		return edge, err
	}

	err = svc.q.AddSuggestNodeTask(ctx, node)
	if err != nil {
		return edge, err
	}

	return edge, nil
}

func (svc *NodeService) BulkCreateNode(ctx context.Context, nodes *[]models.Node) error {
	return svc.node.CreateBulkNodes(ctx, nodes)
}

func NodeServiceBuilder(db *gorm.DB, q *asynq.Client) *NodeService {
	nodeRepo := repository.NodeRepositoryBuilder(db)
	edgeRepo := repository.EdgeRepositoryBuilder(db)
	queue := NewQueueService(q)
	return &NodeService{node: nodeRepo, edge: edgeRepo, q: queue}
}
