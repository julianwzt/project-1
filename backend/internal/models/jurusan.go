package models

type Jurusan struct {
	IDJurusan   int    `json:"id_jurusan"`
	NamaJurusan string `json:"nama_jurusan"`
	Fakultas    string `json:"fakultas"`
	Jenjang     string `json:"jenjang"`
}