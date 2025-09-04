package dto

import "github.com/neilchetna/mind-mapper/internal/models"

type CreateNodeResponse struct {
	Node models.Node `json:"node"`
	Edge models.Edge `json:"edge"`
}
