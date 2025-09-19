package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
	"github.com/neilchetna/mind-mapper/internal/models"
	workerDTO "github.com/neilchetna/mind-mapper/internal/worker/dto"
	"github.com/neilchetna/mind-mapper/pkg"
)

type QueueService struct {
	q *asynq.Client
}

func (svc *QueueService) AddSuggestNodeTask(ctx context.Context, node *models.Node) error {
	payload := workerDTO.SuggestNodesPayload{
		ChartId: node.ChartId.String(),
		NodeId:  node.ID.String(),
		UserId:  node.UserId.String(),
	}

	_, err := svc.AddNewTask(ctx, pkg.Tasks.SuggestNodes, payload)
	return err
}

func (svc *QueueService) AddNewTask(ctx context.Context, taskId string, data any, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(taskId, payload)
	info, err := svc.q.EnqueueContext(ctx, task, opts...)

	if err != nil {
		return nil, err
	}

	log.Print("New Task added")
	return info, nil
}

func NewQueueService(q *asynq.Client) *QueueService {
	return &QueueService{q}
}
