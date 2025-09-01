package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/models"
	"gorm.io/gorm"
)

type NodeRepository struct {
	db *gorm.DB
}

func NodeRepositoryBuilder(db *gorm.DB) *NodeRepository {
	return &NodeRepository{db}
}

func (r *NodeRepository) CreatNode(ctx context.Context, node *models.Node) error {
	res := r.db.WithContext(ctx).Create(node)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *NodeRepository) GetNode(ctx context.Context, chartID uuid.UUID, userID uuid.UUID) ([]models.Node, error) {
	var nodes []models.Node
	res := r.db.WithContext(ctx).Where(&models.Node{ChartId: chartID, UserId: userID}).Find(&nodes)

	if res.Error != nil {
		return nil, res.Error
	}

	return nodes, nil
}
