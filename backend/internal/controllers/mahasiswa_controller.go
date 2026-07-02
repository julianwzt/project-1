package controllers

import (
	"backend-api/internal/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type MahasiswaController struct {
	db *sql.DB
}

func NewMahasiswaController(db *sql.DB) *MahasiswaController {
	return &MahasiswaController{db: db}
}

// @Summary      Mengambil Data Mahasiswa
// @Description  Mengambil semua data mahasiswa beserta detail jurusannya
// @Tags         mahasiswa
// @Produce      json
// @Success      200  {array}   models.Mahasiswa "Daftar data mahasiswa"
// @Router       /mahasiswa [get]
func (a *MahasiswaController) GetMahasiswa(c *gin.Context) {
	query := `
        SELECT m.id, m.nama, m.umur, m.nim, TO_CHAR(m.tgl_lahir, 'YYYY-MM-DD'), m.alamat, m.id_jurusan, 
			COALESCE(j.nama_jurusan, ''), COALESCE(j.fakultas, ''), COALESCE(j.jenjang, '')
        FROM mahasiswa m
        LEFT JOIN jurusan j ON m.id_jurusan = j.id_jurusan`

	rows, err := a.db.Query(query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Mahasiswa
	for rows.Next() {
		var m models.Mahasiswa
		var j models.Jurusan
		rows.Scan(&m.ID, &m.Nama, &m.Umur, &m.NIM, &m.TglLahir, &m.Alamat, &m.IDJurusan, &j.NamaJurusan, &j.Fakultas, &j.Jenjang)
		j.IDJurusan = m.IDJurusan
		m.DetailJurusan = &j
		list = append(list, m)
	}
	c.JSON(200, list)
}

// @Summary      Menambahkan Data Mahasiswa
// @Description  Menambahkan data mahasiswa baru ke database
// @Tags         mahasiswa
// @Accept       json
// @Produce      json
// @Param        mahasiswa  body    models.Mahasiswa  true  "Data Mahasiswa JSON"
// @Success      201  {object}  models.Mahasiswa "Data mahasiswa berhasil disimpan"
// @Failure      400  {object}  map[string]string "Format data tidak valid"
// @Router       /mahasiswa [post]
func (a *MahasiswaController) CreateMahasiswa(c *gin.Context) {
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	query := `INSERT INTO mahasiswa (nama, umur, nim, tgl_lahir, alamat, id_jurusan) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := a.db.QueryRow(query, input.Nama, input.Umur, input.NIM, input.TglLahir, input.Alamat, input.IDJurusan).Scan(&input.ID)
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
// @Success      200  {object}  map[string]string "Data mahasiswa berhasil diperbarui"
// @Failure      404  {object}  map[string]string "Mahasiswa dengan ID tersebut tidak ditemukan"
// @Router       /mahasiswa/{id} [put]
func (a *MahasiswaController) UpdateMahasiswa(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid: " + err.Error()})
		return
	}
	query := `UPDATE mahasiswa SET nama=$1, umur=$2, nim=$3, tgl_lahir=$4, alamat=$5, id_jurusan=$6 WHERE id=$7`
	result, err := a.db.Exec(query, input.Nama, input.Umur, input.NIM, input.TglLahir, input.Alamat, input.IDJurusan, id)
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
// @Success      200  {object}  map[string]string "Data mahasiswa berhasil dihapus"
// @Failure      404  {object}  map[string]string "Mahasiswa dengan ID tersebut tidak ditemukan"
// @Router       /mahasiswa/{id} [delete]
func (a *MahasiswaController) DeleteMahasiswa(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}
	query := `DELETE FROM mahasiswa WHERE id=$1`
	result, err := a.db.Exec(query, id)
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

// ExportMahasiswaExcel godoc
// @Summary      Export Data Mahasiswa ke Excel
// @Description  Mengunduh seluruh daftar data mahasiswa ke dalam file Excel (.xlsx) dengan format tanggal bersih (YYYY-MM-DD).
// @Tags         mahasiswa
// @Produce      application/octet-stream
// @Success      200  {file}    binary "File Laporan_Mahasiswa.xlsx berhasil diunduh"
// @Failure      500  {object}  map[string]interface{} "Gagal memproses data atau menulis file Excel"
// @Router       /mahasiswa/export [get]
func (a *MahasiswaController) ExportMahasiswaExcel(c *gin.Context) {
	query := `
		SELECT m.nama, m.umur, m.nim, TO_CHAR(m.tgl_lahir, 'DD-MM-YYYY'), m.alamat, j.nama_jurusan
		FROM mahasiswa m
		LEFT JOIN jurusan j ON m.id_jurusan = j.id_jurusan
	`
	rows, err := a.db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	f := excelize.NewFile()
	sheetName := "Data Mahasiswa"
	f.SetSheetName("Sheet1", sheetName)

	headers := []string{"No", "NIM", "Nama Lengkap", "Umur", "Tanggal Lahir", "Alamat", "Jurusan"}
	for colIdx, headerText := range headers {
		cell, _ := excelize.CoordinatesToCellName(colIdx+1, 1)
		f.SetCellValue(sheetName, cell, headerText)
	}

	rowNum := 2
	no := 1
	for rows.Next() {
		var nama, nim, tglLahir, alamat string
		var umur int
		var namaJurusan *string

		if err := rows.Scan(&nama, &umur, &nim, &tglLahir, &alamat, &namaJurusan); err == nil {
			f.SetCellValue(sheetName, "A"+strconv.Itoa(rowNum), no)
			f.SetCellStr(sheetName, "B"+strconv.Itoa(rowNum), nim)
			f.SetCellValue(sheetName, "C"+strconv.Itoa(rowNum), nama)
			f.SetCellValue(sheetName, "D"+strconv.Itoa(rowNum), umur)
			f.SetCellStr(sheetName, "E"+strconv.Itoa(rowNum), tglLahir)
			f.SetCellValue(sheetName, "F"+strconv.Itoa(rowNum), alamat)

			if namaJurusan != nil {
				f.SetCellValue(sheetName, "G"+strconv.Itoa(rowNum), *namaJurusan)
			} else {
				f.SetCellValue(sheetName, "G"+strconv.Itoa(rowNum), "-")
			}

			rowNum++
			no++
		}
	}

	c.Header("Content-Disposition", "attachment; filename=Laporan_Mahasiswa.xlsx")
	c.Header("Content-Type", "application/octet-stream")

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat file Excel"})
	}
}
