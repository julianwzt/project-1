package models

type Mahasiswa struct {
	ID            int      `json:"id"`
	Nama          string   `json:"nama"`
	Umur          int      `json:"umur"`
	NIM           string   `json:"nim"`
	TglLahir      string   `json:"tgl_lahir"`
	Alamat        string   `json:"alamat"`
	IDJurusan     int      `json:"id_jurusan"`
	DetailJurusan *Jurusan `json:"jurusan"`
}
