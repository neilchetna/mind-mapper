package models

import (
	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/pkg"
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
	IsSuggested bool      `json:"isSuggested" gorm:"default:false"`
}

func (p *Node) AfterCreate(tx *gorm.DB) error {
	if p.IsSeedNode {
		return nil
	}

	edge := Edge{
		ChartId: p.ChartId,
		Target:  p.ID,
		Source:  p.ParentId,
	}

	if err := tx.Create(&edge).Error; err != nil {
		return err
	}

	return nil
}

func (p *Node) BeforeSave(tx *gorm.DB) (err error) {
	if p.UserId == uuid.Nil {
		return pkg.ErrUserIdNull
	}

	return
}
