package controllers

import (
	"database/sql"
	"net/http"

	"backend-api/internal/models"

	"github.com/gin-gonic/gin"
)

type JurusanController struct {
	DB *sql.DB
}

func NewJurusanController(db *sql.DB) *JurusanController {
	return &JurusanController{DB: db}
}

// @Summary GetJurusan
// @Description Mengambil seluruh data referensi jurusan untuk dropdown
// @Tags Jurusan
// @Accept  json
// @Produce json
// @Success 200 {object} []models.Jurusan "Data jurusan berhasil diambil"
// @Failure 500 {object} map[string]string "Gagal mengambil data jurusan"
// @Router /jurusan [get]
func (jc *JurusanController) GetJurusan(c *gin.Context) {
	rows, err := jc.DB.Query("SELECT id_jurusan, nama_jurusan, fakultas, jenjang FROM Jurusan")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data jurusan: " + err.Error()})
		return
	}
	defer rows.Close()

	var jurusanList []models.Jurusan
	for rows.Next() {
		var jur models.Jurusan
		if err := rows.Scan(&jur.IDJurusan, &jur.NamaJurusan, &jur.Fakultas, &jur.Jenjang); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data: " + err.Error()})
			return
		}
		jurusanList = append(jurusanList, jur)
	}

	if jurusanList == nil {
		jurusanList = []models.Jurusan{}
	}

	c.JSON(http.StatusOK, jurusanList)
}
