package worker

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	llmDTO "github.com/neilchetna/mind-mapper/internal/llm/dto"
	"github.com/neilchetna/mind-mapper/internal/models"
	"github.com/neilchetna/mind-mapper/internal/service"
	workerDTO "github.com/neilchetna/mind-mapper/internal/worker/dto"
	"gorm.io/gorm"
)

type SuggestNodesWorker struct {
	node *service.NodeService
	edge *service.EdgeService
	llm  *service.LLMService
}

func (w *SuggestNodesWorker) SuggestNodesHandler(ctx context.Context, task *asynq.Task) error {
	LoggerStart(ctx, task)
	payloadString := task.Payload()

	var payload workerDTO.SuggestNodesPayload
	err := json.Unmarshal(payloadString, &payload)

	if err != nil {
		LoggerErr(ctx, task, err)
		return err
	}

	chartId, err := uuid.Parse(payload.ChartId)
	if err != nil {
		LoggerErr(ctx, task, err)
		return err
	}

	nodeId, err := uuid.Parse(payload.NodeId)
	if err != nil {
		LoggerErr(ctx, task, err)
		return err
	}

	userId, err := uuid.Parse(payload.UserId)

	nodes, err := w.node.GetChartNodes(ctx, chartId)
	if err != nil {
		LoggerErr(ctx, task, err)
		return err
	}

	edges, err := w.edge.GetEdges(ctx, chartId)
	if err != nil {
		LoggerErr(ctx, task, err)
		return err
	}

	res, err := w.llm.GenerateNodes(nodes, edges, nodeId)
	if err != nil {
		LoggerErr(ctx, task, err)
		return err
	}

	newNodes, err := tokenizeNodesFromResponse(res, nodeId, userId, chartId)
	if err != nil {
		LoggerErr(ctx, task, err)
		return err
	}

	if err = w.node.BulkCreateNode(ctx, &newNodes); err != nil {
		LoggerErr(ctx, task, err)
		return err
	}

	LoggerSuccess(ctx, task)
	return nil
}

func tokenizeNodesFromResponse(res string, parentId uuid.UUID, userId uuid.UUID, chartId uuid.UUID) ([]models.Node, error) {
	var parsedRes llmDTO.SuggestNodesResponse
	err := json.Unmarshal([]byte(res), &parsedRes)
	if err != nil {
		return parsedRes.Nodes, err
	}

	var nodes []models.Node
	textList := parsedRes.Nodes

	for _, text := range textList {
		newNode := models.Node{ParentId: parentId, UserId: userId, ChartId: chartId, Text: text.Text, IsSuggested: true}
		nodes = append(nodes, newNode)
	}

	return nodes, nil
}

func SuggestNodesWorkerBuilder(db *gorm.DB) *SuggestNodesWorker {
	node := service.NodeServiceBuilder(db, nil)
	edge := service.NewEdgeServiceBuilder(db)
	llm := &service.LLMService{}

	return &SuggestNodesWorker{node, edge, llm}
}
