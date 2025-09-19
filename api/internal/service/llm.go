package service

import (
	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/internal/llm"
	"github.com/neilchetna/mind-mapper/internal/llm/prompts"
	"github.com/neilchetna/mind-mapper/internal/models"
)

type LLMService struct {
}

func (svc *LLMService) GenerateNodes(nodes []models.Node, edges []models.Edge, targetNodeID uuid.UUID) (string, error) {
	textGraph := llm.BuildPromptFromGraph(nodes, edges, targetNodeID)

	prompt, err := prompts.SuggestNodesPrompt(textGraph)
	if err != nil {
		return "", err
	}

	llmClient := &llm.LLMClient{}
	res, err := llmClient.MakeLLMCall(prompt)
	if err != nil {
		return "", err
	}

	return res, nil
}
