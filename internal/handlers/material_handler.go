package handlers

import (
	"net/http"
	"strconv"

	"challecara2025-back/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MaterialHandler struct {
	db *gorm.DB
}

func NewMaterialHandler(db *gorm.DB) *MaterialHandler {
	return &MaterialHandler{db: db}
}

type materialCreateInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type materialUpdateInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateMaterial 新しい参考資料を作成
func (h *MaterialHandler) CreateMaterial(c *gin.Context) {
	bookIDParam := c.Param("id")
	bookIDUint64, err := strconv.ParseUint(bookIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var input materialCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 資料が紐づくBookの存在確認
	var book models.Book
	if err := h.db.First(&book, bookIDParam).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify book"})
		return
	}

	material := models.Material{
		BookID:  uint(bookIDUint64),
		Title:   input.Title,
		Content: input.Content,
	}

	if err := h.db.Create(&material).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create material"})
		return
	}

	c.JSON(http.StatusCreated, material)
}

// GetMaterials 特定のBookに紐づく参考資料を取得
func (h *MaterialHandler) GetMaterials(c *gin.Context) {
	bookIDParam := c.Param("id")

	if _, err := strconv.ParseUint(bookIDParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := h.db.First(&models.Book{}, bookIDParam).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify book"})
		return
	}

	var materials []models.Material
	if err := h.db.Where("book_id = ?", bookIDParam).Order("created_at DESC").Find(&materials).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch materials"})
		return
	}

	c.JSON(http.StatusOK, materials)
}

// GetMaterial 特定の参考資料を取得
func (h *MaterialHandler) GetMaterial(c *gin.Context) {
	id := c.Param("id")
	var material models.Material

	if err := h.db.First(&material, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch material"})
		return
	}

	c.JSON(http.StatusOK, material)
}

// UpdateMaterial 参考資料を更新
func (h *MaterialHandler) UpdateMaterial(c *gin.Context) {
	id := c.Param("id")
	var material models.Material

	if err := h.db.First(&material, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch material"})
		return
	}

	var input materialUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	material.Title = input.Title
	material.Content = input.Content

	if err := h.db.Save(&material).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update material"})
		return
	}

	c.JSON(http.StatusOK, material)
}

// DeleteMaterial 参考資料を削除
func (h *MaterialHandler) DeleteMaterial(c *gin.Context) {
	id := c.Param("id")

	if err := h.db.Delete(&models.Material{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete material"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}
