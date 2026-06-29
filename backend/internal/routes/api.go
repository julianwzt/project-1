package routes

import (
	"backend-api/internal/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	mahasiswaController := controllers.NewMahasiswaController(db)
	jurusanController := controllers.NewJurusanController(db)

	r.GET("/api/mahasiswa", mahasiswaController.GetMahasiswa)
	r.POST("/api/mahasiswa", mahasiswaController.CreateMahasiswa)
	r.PUT("/api/mahasiswa/:id", mahasiswaController.UpdateMahasiswa)
	r.DELETE("/api/mahasiswa/:id", mahasiswaController.DeleteMahasiswa)
	r.GET("/api/mahasiswa/export", mahasiswaController.ExportMahasiswaExcel)
	r.GET("/api/jurusan", jurusanController.GetJurusan)
}
