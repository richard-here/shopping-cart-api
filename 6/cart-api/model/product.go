package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Code      uuid.UUID      `gorm:"type:uuid;primaryKey" json:"code"`
	Name      string         `json:"name" validate:"required"`
	Quantity  int            `json:"qty" validate:"required,min=0,numeric"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Products struct {
	Products []Products `json:"products"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Code == uuid.Nil {
		p.Code = uuid.New()
	}
	return
}
