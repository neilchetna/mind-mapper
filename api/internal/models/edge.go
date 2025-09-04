package models

import "github.com/google/uuid"

type Edge struct {
	BaseModel
	Source  uuid.UUID `json:"source" gorm:"type:uuid"`
	Target  uuid.UUID `json:"target" gorm:"type:uuid"`
	ChartId uuid.UUID `json:"chartId" gorm:"type:uuid"`
}
