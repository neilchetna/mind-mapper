package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chart struct {
	BaseModel
	Title string `json:"title" gorm:"default:New map"`
	Description string `json:"description"`
	UserId uuid.UUID `json:"userId" gorm:"type:uuid"`
}


func (p *Chart) BeforeSave(tx *gorm.DB) (err error) {
	if p.UserId == uuid.Nil {
		return ErrUserIdNull
	}

	return
}