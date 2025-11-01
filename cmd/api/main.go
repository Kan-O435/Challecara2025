package main

import (
	"log"
	"os"

	"challecara2025-back/internal/database"
	"challecara2025-back/internal/handlers"
	"challecara2025-back/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// データベースに接続
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// マイグレーション実行
	if err := database.Migrate(&models.Book{}, &models.Episode{}, &models.Material{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Ginルーターを初期化
	router := gin.Default()

	// CORSミドルウェアを追加
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// ハンドラーを初期化
	db := database.GetDB()
	bookHandler := handlers.NewBookHandler(db)
	episodeHandler := handlers.NewEpisodeHandler(db)
	materialHandler := handlers.NewMaterialHandler(db)

	// APIルートを設定
	api := router.Group("/api")
	{
		// 資料関連のルート
		books := api.Group("/books")
		{
			books.POST("", bookHandler.CreateBook)
			books.GET("", bookHandler.GetBooks)
			books.GET("/:id", bookHandler.GetBook)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)

			// エピソード関連のルート（資料配下）- パラメータ名を :id に統一
			books.POST("/:id/episodes", episodeHandler.CreateEpisode)
			books.GET("/:id/episodes", episodeHandler.GetEpisodes)

			// 参考資料関連のルート（資料配下）
			books.POST("/:id/materials", materialHandler.CreateMaterial)
			books.GET("/:id/materials", materialHandler.GetMaterials)
		}

		// エピソード関連のルート（直接アクセス）
		episodes := api.Group("/episodes")
		{
			episodes.GET("/:id", episodeHandler.GetEpisode)
			episodes.PUT("/:id", episodeHandler.UpdateEpisode)
			episodes.DELETE("/:id", episodeHandler.DeleteEpisode)
		}

		// 参考資料関連のルート（直接アクセス）
		materials := api.Group("/materials")
		{
			materials.GET("/:id", materialHandler.GetMaterial)
			materials.PUT("/:id", materialHandler.UpdateMaterial)
			materials.DELETE("/:id", materialHandler.DeleteMaterial)
		}
	}

	// ヘルスチェック用エンドポイント
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// サーバーを起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
