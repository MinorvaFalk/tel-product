package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Products struct {
	ID          uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Price       float64         `json:"price"`
	Variety     *string         `json:"variety"`
	Rating      float32         `json:"rating"`
	Stock       int             `json:"stock"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`
}

func (Products) TableName() string {
	return "products"
}
