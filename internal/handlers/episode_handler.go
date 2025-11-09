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
	bookID := c.Param("id")
	var episode models.Episode

	if err := c.ShouldBindJSON(&episode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// book_idをパラメータから設定
	var bookIDUint uint64
	bookIDUint, err := strconv.ParseUint(bookID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	episode.BookID = uint(bookIDUint)

	// 資料が存在するか確認
	var book models.Book
	if err := h.db.First(&book, bookID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify book"})
		return
	}

	if err := h.db.Create(&episode).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create episode"})
		return
	}

	c.JSON(http.StatusCreated, episode)
}

// GetEpisodes 特定の資料のすべてのエピソードを取得
func (h *EpisodeHandler) GetEpisodes(c *gin.Context) {
	bookID := c.Param("id")
	var episodes []models.Episode

	if err := h.db.Where("book_id = ?", bookID).Order("episode_no").Find(&episodes).Error; err != nil {
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
