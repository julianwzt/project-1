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
    ('Budi Santoso', 21, '13519001', '2003-05-10', 'Jl. Merdeka No. 10', 1),
    ('Andi Setiawan', 20, '13519002', '2006-03-15', 'Jl. Sudirman No. 45, Jakarta', 2),
    ('Citra Lestari', 22, '13519003', '2004-11-20', 'Jl. Asia Afrika No. 12, Bandung', 3),
    ('Dewi Sartika', 19, '13519004', '2007-01-05', 'Jl. Pahlawan No. 88, Semarang', 1),
    ('Eko Prasetyo', 21, '13519005', '2005-08-30', 'Jl. Gajah Mada No. 101, Surabaya', 4),
    ('Fajar Nugroho', 20, '13519006', '2006-02-14', 'Jl. Malioboro No. 5, Yogyakarta', 5),  
    ('Gita Savitri', 21, '13519007', '2005-09-25', 'Jl. Dago No. 22, Bandung', 6),
    ('Hadi Sucipto', 22, '13519008', '2004-12-12', 'Jl. Pemuda No. 7, Medan', 7),
    ('Intan Permatasari', 19, '13519009', '2007-04-18', 'Jl. Gatot Subroto No. 33, Jakarta', 1),
    ('Joko Anwar', 23, '13519010', '2003-06-21', 'Jl. Slamet Riyadi No. 99, Solo', 2),
    ('Kartika Putri', 20, '13519011', '2006-10-10', 'Jl. Veteran No. 44, Malang', 3);