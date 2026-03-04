package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:char(36);primarykey" json:"id"`
	Username     string         `gorm:"size:100;uniqueIndex;not null" json:"username"`
	Email        string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	PasswordHash string         `gorm:"size:255;not null" json:"-"` // JSONには含めない
	DisplayName  string         `gorm:"size:255" json:"display_name"`
	Bio          string         `gorm:"type:text" json:"bio"`
	AvatarURL    string         `gorm:"size:500" json:"avatar_url,omitempty"`
	Books        []Book         `gorm:"foreignKey:AuthorID" json:"books,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
