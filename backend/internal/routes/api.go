package routes

import (
	"backend-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/mahasiswa", controllers.GetMahasiswa)
	r.POST("/api/mahasiswa", controllers.CreateMahasiswa)
	r.PUT("/api/mahasiswa/{id}", controllers.UpdateMahasiswa)
	r.DELETE("/api/mahasiswa/{id}", controllers.DeleteMahasiswa)
}