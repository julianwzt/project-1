package models

type Jurusan struct {
	IDJurusan   int    `json:"id_jurusan" binding:"required"`
	NamaJurusan string `json:"nama_jurusan" binding:"required"`
	Fakultas    string `json:"fakultas" binding:"required"`
	Jenjang     string `json:"jenjang" binding:"required" validate:"oneof=S1 D3 D4 S2 S3"`
}

type Mahasiswa struct {
	ID            int      `json:"id"`
	Nama          string   `json:"nama" binding:"required"`
	Umur          int      `json:"umur" binding:"required"`
	NIM           string   `json:"nim" binding:"required,min=12,max=12"`
	TglLahir      string   `json:"tgl_lahir" binding:"required"`
	Alamat        string   `json:"alamat" binding:"required"`
	IDJurusan     int      `json:"id_jurusan" binding:"required"`
	DetailJurusan *Jurusan `json:"jurusan"`
}
