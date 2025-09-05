package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Node struct {
	BaseModel
	Text        string    `json:"text" validate:"required" gorm:"not null"`
	Description string    `json:"description"`
	IsSeedNode  bool      `json:"isSeedNode" validate:"required"`
	UserId      uuid.UUID `json:"userId" gorm:"type:uuid"`
	ChartId     uuid.UUID `json:"chartId" gorm:"type:uuid"`
	ParentId    uuid.UUID `json:"parentId" gorm:"type:uuid"`
}

func (p *Node) BeforeSave(tx *gorm.DB) (err error) {
	if p.UserId == uuid.Nil {
		return ErrUserIdNull
	}

	return
}
