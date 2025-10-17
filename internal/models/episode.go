package models

import (
	"time"
	"gorm.io/gorm"
)

type Episode struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	NovelID   uint           `gorm:"not null;index" json:"novel_id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Content   string         `gorm:"type:longtext;not null" json:"content"`
	EpisodeNo int            `gorm:"not null" json:"episode_no"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}