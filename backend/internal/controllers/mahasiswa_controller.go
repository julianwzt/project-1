package controllers

import (
	"net/http"

	"backend-api/internal/config"
	"backend-api/internal/models"

	"github.com/gin-gonic/gin"
)

// @Summary      Mengambil Data Mahasiswa
// @Description  Mengambil semua data mahasiswa beserta detail jurusannya
// @Tags         mahasiswa
// @Produce      json
// @Success      200  {array}   models.Mahasiswa
// @Router       /api/mahasiswa [get]
func GetMahasiswa(c *gin.Context) {
	query := `
		SELECT m.id, m.nama, m.umur, m.nim, m.tgl_lahir, m.alamat, m.id_jurusan, j.nama_jurusan, j.fakultas, j.jenjang
		FROM mahasiswa m
		LEFT JOIN jurusan j ON m.id_jurusan = j.id_jurusan
	`
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var listMahasiswa []models.Mahasiswa
	for rows.Next() {
		var m models.Mahasiswa
		var j models.Jurusan
		err := rows.Scan(&m.ID, &m.Nama, &m.Umur, &m.NIM, &m.TglLahir, &m.Alamat, &m.IDJurusan, &j.NamaJurusan, &j.Fakultas, &j.Jenjang)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		j.IDJurusan = m.IDJurusan
		m.DetailJurusan = &j
		listMahasiswa = append(listMahasiswa, m)
	}
	c.JSON(http.StatusOK, listMahasiswa)
}

// @Summary      Menambahkan Data Mahasiswa
// @Description  Menambahkan data mahasiswa baru ke database
// @Tags         mahasiswa
// @Accept       json
// @Produce      json
// @Param        mahasiswa  body    models.Mahasiswa  true  "Data Mahasiswa JSON"
// @Success      201  {object}  models.Mahasiswa
// @Router       /api/mahasiswa [post]
func CreateMahasiswa(c *gin.Context) {
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	query := `INSERT INTO mahasiswa (nama, umur, nim, tgl_lahir, alamat, id_jurusan) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := config.DB.QueryRow(query, input.Nama, input.Umur, input.NIM, input.TglLahir, input.Alamat, input.IDJurusan).Scan(&input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Data mahasiswa berhasil disimpan", "data": input})
}

// @Summary      Mengupdate Data Mahasiswa
// @Description  Mengupdate data mahasiswa berdasarkan ID
// @Tags         mahasiswa
// @Accept       json
// @Produce      json
// @Param        id   path      int               true  "ID Mahasiswa"
// @Param        mahasiswa  body    models.Mahasiswa  true  "Data Mahasiswa JSON"
// @Success      200  {object}  map[string]string
// @Router       /api/mahasiswa/{id} [put]
func UpdateMahasiswa(c *gin.Context) {
	id := c.Param("id")
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	query := `UPDATE mahasiswa SET nama=$1, umur=$2, nim=$3, tgl_lahir=$4, alamat=$5, id_jurusan=$6 WHERE id=$7`
	result, err := config.DB.Exec(query, input.Nama, input.Umur, input.NIM, input.TglLahir, input.Alamat, input.IDJurusan, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data: " + err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa dengan ID tersebut tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data mahasiswa berhasil diperbarui"})
}

// @Summary      Menghapus Data Mahasiswa
// @Description  Menghapus data mahasiswa berdasarkan ID
// @Tags         mahasiswa
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID Mahasiswa"
// @Success      200  {object}  map[string]string
// @Router       /api/mahasiswa/{id} [delete]
func DeleteMahasiswa(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM mahasiswa WHERE id=$1`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data: " + err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa dengan ID tersebut tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data mahasiswa berhasil dihapus"})
}
