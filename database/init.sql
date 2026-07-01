CREATE TABLE IF NOT EXISTS Jurusan (
    id_jurusan SERIAL PRIMARY KEY,
    nama_jurusan VARCHAR(100) NOT NULL,
    fakultas VARCHAR(100) NOT NULL,
    jenjang VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS Mahasiswa (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(150) NOT NULL,
    umur INT NOT NULL,
    nim VARCHAR(20) UNIQUE NOT NULL,
    tgl_lahir DATE NOT NULL,
    alamat TEXT NOT NULL,
    id_jurusan INT NOT NULL,
    CONSTRAINT fk_jurusan FOREIGN KEY(id_jurusan) REFERENCES Jurusan(id_jurusan) ON DELETE RESTRICT
);

INSERT INTO Jurusan (nama_jurusan, fakultas, jenjang) VALUES
    ('Teknik Informatika', 'Fakultas Informatika', 'S1'),
    ('Sistem Informasi', 'Fakultas Informatika', 'S1'),
    ('Teknik Elektro', 'Fakultas Teknik', 'S1'),
    ('Teknik Mesin', 'Fakultas Teknik', 'S1'),
    ('Desain Komunikasi Visual', 'Fakultas Industri Kreatif', 'D4'),
    ('Manajemen', 'Fakultas Ekonomi dan Bisnis', 'S1'),
    ('Akuntansi', 'Fakultas Ekonomi dan Bisnis', 'S1');

    INSERT INTO Mahasiswa (nama, umur, nim, tgl_lahir, alamat, id_jurusan) VALUES
    ('Budi Santoso', 21, '13519001', '2003-05-10', 'Jl. Merdeka No. 10', 1);