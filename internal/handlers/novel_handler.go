package handlers

import (
	"net/http"

	"challecara2025-back/internal/models"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NovelHandler struct {
	db *gorm.DB
}

func NewNovelHandler(db *gorm.DB) *NovelHandler {
	return &NovelHandler{db: db}
}

// CreateNovel 新しい小説を作成
func (h *NovelHandler) CreateNovel(c *gin.Context) {
	var novel models.Novel
	
	if err := c.ShouldBindJSON(&novel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&novel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create novel"})
		return
	}

	c.JSON(http.StatusCreated, novel)
}

// GetNovels すべての小説を取得
func (h *NovelHandler) GetNovels(c *gin.Context) {
	var novels []models.Novel
	
	if err := h.db.Preload("Episodes").Find(&novels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch novels"})
		return
	}

	c.JSON(http.StatusOK, novels)
}

// GetNovel 特定の小説を取得
func (h *NovelHandler) GetNovel(c *gin.Context) {
	id := c.Param("id")
	var novel models.Novel

	if err := h.db.Preload("Episodes").First(&novel, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Novel not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch novel"})
		return
	}

	c.JSON(http.StatusOK, novel)
}

// UpdateNovel 小説を更新
func (h *NovelHandler) UpdateNovel(c *gin.Context) {
	id := c.Param("id")
	var novel models.Novel

	if err := h.db.First(&novel, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Novel not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch novel"})
		return
	}

	if err := c.ShouldBindJSON(&novel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Save(&novel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update novel"})
		return
	}

	c.JSON(http.StatusOK, novel)
}

// DeleteNovel 小説を削除
func (h *NovelHandler) DeleteNovel(c *gin.Context) {
	id := c.Param("id")
	
	if err := h.db.Delete(&models.Novel{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete novel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Novel deleted successfully"})
}