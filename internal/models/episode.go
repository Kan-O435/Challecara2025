package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Episode struct {
	ID        uuid.UUID      `gorm:"type:char(36);primarykey" json:"id"`
	BookID    uuid.UUID      `gorm:"type:char(36);not null;index" json:"book_id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Content   string         `gorm:"type:longtext;not null" json:"content"`
	EpisodeNo int            `gorm:"not null" json:"episode_no"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
