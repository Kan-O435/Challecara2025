package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Material struct {
	ID        uuid.UUID      `gorm:"type:char(36);primarykey" json:"id"`
	BookID    uuid.UUID      `gorm:"type:char(36);not null;index" json:"book_id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Content   string         `gorm:"type:longtext;not null" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
