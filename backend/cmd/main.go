package main

import (
	"net/http"

	_ "backend-api/docs"
	"backend-api/internal/config"
	"backend-api/internal/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Manajemen Mahasiswa
// @version         1.0
// @description     Ini adalah dokumentasi REST API untuk mengelola data mahasiswa dan jurusan.
// @host            localhost:8080
// @BasePath        /api
// @Schemes         http
func main() {
	config.ConnectDB()
	defer config.DB.Close()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, EXPORT, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r, config.DB)
	r.Run(":8080")
}
