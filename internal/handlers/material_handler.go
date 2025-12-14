package handlers

import (
	"net/http"

	"challecara2025-back/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	bookUUID, err := uuid.Parse(bookIDParam)
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
	if err := h.db.Where("id = ?", bookUUID).First(&book).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify book"})
		return
	}

	// Generate UUIDv7 for the new material
	newID, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate UUID"})
		return
	}

	material := models.Material{
		ID:      newID,
		BookID:  bookUUID,
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

	bookUUID, err := uuid.Parse(bookIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := h.db.Where("id = ?", bookUUID).First(&models.Book{}).Error; err != nil {
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

// GetMaterialsByIDs 複数の資料をID指定で取得
func (h *MaterialHandler) GetMaterialsByIDs(c *gin.Context) {
	bookIDParam := c.Param("id")

	var input struct {
		IDs []uuid.UUID `json:"ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var materials []models.Material
	query := h.db.Where("id IN ?", input.IDs)

	// book_idがパスに含まれている場合はフィルタリング
	if bookIDParam != "" {
		bookID, err := uuid.Parse(bookIDParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}
		query = query.Where("book_id = ?", bookID)
	}

	if err := query.Find(&materials).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch materials"})
		return
	}

	c.JSON(http.StatusOK, materials)
}

// GetMaterial 特定の参考資料を取得
func (h *MaterialHandler) GetMaterial(c *gin.Context) {
	id := c.Param("id")
	var material models.Material

	materialID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	if err := h.db.Where("id = ?", materialID).First(&material).Error; err != nil {
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

	materialID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	if err := h.db.Where("id = ?", materialID).First(&material).Error; err != nil {
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

	materialID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	if err := h.db.Where("id = ?", materialID).Delete(&models.Material{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete material"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}
