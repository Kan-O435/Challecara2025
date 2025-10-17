package handlers

import (
	"net/http"
	"strconv"

	"challecara2025-back/internal/models"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EpisodeHandler struct {
	db *gorm.DB
}

func NewEpisodeHandler(db *gorm.DB) *EpisodeHandler {
	return &EpisodeHandler{db: db}
}

// CreateEpisode 新しいエピソードを作成
func (h *EpisodeHandler) CreateEpisode(c *gin.Context) {
	novelID := c.Param("id")
	var episode models.Episode
	
	if err := c.ShouldBindJSON(&episode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// novel_idをパラメータから設定
	var novelIDUint uint64
	novelIDUint, err := strconv.ParseUint(novelID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid novel ID"})
		return
	}
	episode.NovelID = uint(novelIDUint)

	// 小説が存在するか確認
	var novel models.Novel
	if err := h.db.First(&novel, novelID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Novel not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify novel"})
		return
	}

	if err := h.db.Create(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create episode"})
		return
	}

	c.JSON(http.StatusCreated, episode)
}

// GetEpisodes 特定の小説のすべてのエピソードを取得
func (h *EpisodeHandler) GetEpisodes(c *gin.Context) {
	novelID := c.Param("id")
	var episodes []models.Episode
	
	if err := h.db.Where("id = ?", novelID).Order("episode_no").Find(&episodes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch episodes"})
		return
	}

	c.JSON(http.StatusOK, episodes)
}

// GetEpisode 特定のエピソードを取得
func (h *EpisodeHandler) GetEpisode(c *gin.Context) {
	id := c.Param("id")
	var episode models.Episode

	if err := h.db.First(&episode, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Episode not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch episode"})
		return
	}

	c.JSON(http.StatusOK, episode)
}

// UpdateEpisode エピソードを更新
func (h *EpisodeHandler) UpdateEpisode(c *gin.Context) {
	id := c.Param("id")
	var episode models.Episode

	if err := h.db.First(&episode, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Episode not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch episode"})
		return
	}

	if err := c.ShouldBindJSON(&episode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Save(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update episode"})
		return
	}

	c.JSON(http.StatusOK, episode)
}

// DeleteEpisode エピソードを削除
func (h *EpisodeHandler) DeleteEpisode(c *gin.Context) {
	id := c.Param("id")
	
	if err := h.db.Delete(&models.Episode{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete episode"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Episode deleted successfully"})
}