package handlers

import (
	"net/http"

	"challecara2025-back/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
	db *gorm.DB
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{db: db}
}

// CreateBook 新しい資料を作成
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBooks すべての資料を取得
func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []models.Book

	if err := h.db.Preload("Episodes").Preload("Materials").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// GetBook 特定の資料を取得
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := h.db.Preload("Episodes").Preload("Materials").First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook 資料を更新
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := h.db.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook 資料を削除
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := h.db.Delete(&models.Book{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
