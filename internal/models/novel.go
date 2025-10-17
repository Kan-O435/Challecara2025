package models

import (
	"time"
	"gorm.io/gorm"
)

type Novel struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	AuthorID    uint           `json:"author_id"` // 後で認証実装時に使用
	CoverImage  string         `gorm:"size:500" json:"cover_image,omitempty"`
	Genre       string         `gorm:"size:100" json:"genre"`
	Status      string         `gorm:"size:50;default:'draft'" json:"status"` // draft, published, completed
	Episodes    []Episode      `gorm:"foreignKey:NovelID" json:"episodes,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}