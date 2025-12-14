package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID          uuid.UUID      `gorm:"type:char(36);primarykey" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	AuthorID    uuid.UUID      `gorm:"type:char(36)" json:"author_id"` // 認証実装時に使用予定
	CoverImage  string         `gorm:"size:500" json:"cover_image,omitempty"`
	Genre       string         `gorm:"size:100" json:"genre"`
	Status      string         `gorm:"size:50;default:'draft'" json:"status"` // draft, published, completed
	Episodes    []Episode      `gorm:"foreignKey:BookID" json:"episodes,omitempty"`
	Materials   []Material     `gorm:"foreignKey:BookID" json:"materials,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}