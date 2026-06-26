package models

type Jurusan struct {
	IDJurusan   int    `json:"id_jurusan"`
	NamaJurusan string `json:"nama_jurusan"`
	Fakultas    string `json:"fakultas"`
	Jenjang     string `json:"jenjang"`
}

type Mahasiswa struct {
	ID            int      `json:"id"`
	Nama          string   `json:"nama"`
	Umur          int      `json:"umur"`
	NIM           string   `json:"nim"`
	TglLahir      string   `json:"tgl_lahir"`
	Alamat        string   `json:"alamat"`
	IDJurusan     int      `json:"id_jurusan"`
	DetailJurusan *Jurusan `json:"detail_jurusan"`
}
