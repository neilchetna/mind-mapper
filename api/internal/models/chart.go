package models

import (
	"github.com/google/uuid"
	"github.com/neilchetna/mind-mapper/pkg"
	"gorm.io/gorm"
)

type Chart struct {
	BaseModel
	Title       string    `json:"title" gorm:"default:New map"`
	Description string    `json:"description"`
	UserId      uuid.UUID `json:"userId" gorm:"type:uuid"`
	Nodes       []Node    `json:"nodes" validate:"dive,required" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *Chart) BeforeSave(tx *gorm.DB) (err error) {
	if p.UserId == uuid.Nil {
		return pkg.ErrUserIdNull
	}

	return
}
