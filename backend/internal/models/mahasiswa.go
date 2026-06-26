package models

type Jurusan struct {
	IDJurusan   int    `json:"id_jurusan"`
	NamaJurusan string `json:"nama_jurusan"`
	Fakultas    string `json:"fakultas"`
	Jenjang     string `json:"jenjang"`
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
