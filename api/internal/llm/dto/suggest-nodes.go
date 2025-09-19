package dto

import "github.com/neilchetna/mind-mapper/internal/models"

type SuggestNodesResponse struct {
	Nodes []models.Node `json:"nodes"`
}
