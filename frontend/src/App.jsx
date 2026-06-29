import { useState, useEffect } from "react";
import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_URL || "";
const API_URL = `${API_BASE_URL}/api/mahasiswa`;
const API_JURUSAN_URL = `${API_BASE_URL}/api/jurusan`;

const jurusanData = {
  1: { fakultas: "Fakultas Informatika", jenjang: "S1" },
  2: { fakultas: "Fakultas Rekayasa Industri", jenjang: "S1" },
  3: { fakultas: "Fakultas Ilmu Terapan", jenjang: "D3" },
};

function App() {
  const [mahasiswa, setMahasiswa] = useState([]);
  const [searchQuery, setSearchQuery] = useState("");
  const [isEdit, setIsEdit] = useState(false);
  const [selectedId, setSelectedId] = useState(null);
  const [jurusanList, setJurusanList] = useState([]);
  const [formData, setFormData] = useState({
    nama: "",
    umur: "",
    nim: "",
    tgl_lahir: "",
    alamat: "",
    id_jurusan: "",
    fakultas: "",
    jenjang: "",
  });

  const fakultasOptions = Array.from(
    new Set(jurusanList.map((jur) => jur.fakultas)),
  ).filter(Boolean);

  const jenjangOptions = Array.from(
    new Set(jurusanList.map((jur) => jur.jenjang)),
  ).filter(Boolean);

  useEffect(() => {
    fetchMahasiswa();
    fetchJurusan();
  }, []);

  const fetchMahasiswa = async () => {
    try {
      const res = await axios.get(API_URL);
      setMahasiswa(res.data);
    } catch (err) {
      console.error("Gagal mengambil data:", err);
    }
  };

  const fetchJurusan = async () => {
    try {
      const response = await axios.get(API_JURUSAN_URL);
      setJurusanList(response.data);
    } catch (error) {
      console.error("Error fetching jurusan:", error);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    const newForm = { ...formData, [name]: value };

    if (name === "id_jurusan") {
      const selectedJurusan = jurusanList.find(
        (jur) => String(jur.id_jurusan) === value,
      );
      if (selectedJurusan) {
        newForm.fakultas = selectedJurusan.fakultas;
        newForm.jenjang = selectedJurusan.jenjang;
      } else {
        newForm.fakultas = "";
        newForm.jenjang = "";
      }
    }

    if (name === "tgl_lahir" && value) {
      const birthDate = new Date(value);
      const today = new Date();
      let calculatedAge = today.getFullYear() - birthDate.getFullYear();

      const monthDiff = today.getMonth() - birthDate.getMonth();
      if (
        monthDiff < 0 ||
        (monthDiff === 0 && today.getDate() < birthDate.getDate())
      ) {
        calculatedAge--;
      }

      newForm.umur = calculatedAge > 0 ? calculatedAge : 0;
    }

    setFormData(newForm);
  };

  const handleReset = () => {
    setFormData({
      nama: "",
      umur: "",
      nim: "",
      tgl_lahir: "",
      alamat: "",
      id_jurusan: "",
      fakultas: "",
      jenjang: "",
    });
    setIsEdit(false);
    setSelectedId(null);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = {
      nama: formData.nama,
      umur: parseInt(formData.umur),
      nim: formData.nim,
      tgl_lahir: formData.tgl_lahir
        ? new Date(formData.tgl_lahir).toISOString()
        : "",
      alamat: formData.alamat,
      id_jurusan: parseInt(formData.id_jurusan),
    };

    try {
      if (isEdit) {
        await axios.put(`${API_URL}/${selectedId}`, payload);
        alert("Data berhasil diperbarui!");
      } else {
        await axios.post(API_URL, payload);
        alert("Data berhasil disimpan!");
      }
      fetchMahasiswa();
      handleReset();
    } catch (err) {
      alert("Gagal memproses data: " + err.response.data.error);
    }
  };

  const handleDelete = async () => {
    if (!selectedId)
      return alert("Pilih data dari list di sebelah kanan terlebih dahulu!");
    if (window.confirm(`Yakin ingin menghapus data dengan ID ${selectedId}?`)) {
      try {
        await axios.delete(`${API_URL}/${selectedId}`);
        alert("Data berhasil dihapus!");
        fetchMahasiswa();
        handleReset();
      } catch (err) {
        console.error(err);
      }
    }
  };

  const handleItemClick = (m) => {
    setIsEdit(true);
    setSelectedId(m.id);
    const dateOnly = m.tgl_lahir ? m.tgl_lahir.split("T")[0] : "";
    setFormData({
      nama: m.nama,
      umur: m.umur,
      nim: m.nim,
      tgl_lahir: dateOnly,
      alamat: m.alamat,
      id_jurusan: String(m.id_jurusan),
      fakultas: m.jurusan ? m.jurusan.fakultas : "",
      jenjang: m.jurusan ? m.jurusan.jenjang : "",
    });
  };

  const filteredMahasiswa = (mahasiswa || []).filter(
    (m) =>
      m.nama.toLowerCase().includes(searchQuery.toLowerCase()) ||
      m.nim.includes(searchQuery),
  );

  const handleExportExcel = () => {
    window.open(`${API_URL}/export`, "_blank");
  };

  return (
    <div style={styles.container}>
      <div style={styles.leftPanel}>
        <div style={styles.headerForm}>Data Mahasiswa</div>
        <form onSubmit={handleSubmit} style={styles.form}>
          <div style={styles.formGroup}>
            <label style={styles.label}>Nama :</label>
            <input
              type="text"
              name="nama"
              value={formData.nama}
              onChange={handleChange}
              required
              style={styles.input}
            />
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>Umur :</label>
            <input
              type="number"
              name="umur"
              value={formData.umur}
              onChange={handleChange}
              required
              style={styles.input}
            />
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>NIM :</label>
            <input
              type="text"
              name="nim"
              value={formData.nim}
              onChange={handleChange}
              required
              style={styles.input}
            />
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>Tgl Lahir :</label>
            <input
              type="date"
              name="tgl_lahir"
              value={formData.tgl_lahir}
              onChange={handleChange}
              required
              style={styles.input}
            />
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>Alamat :</label>
            <textarea
              name="alamat"
              value={formData.alamat}
              onChange={handleChange}
              rows="3"
              style={styles.textarea}
            />
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>Jurusan :</label>
            <select
              name="id_jurusan"
              value={formData.id_jurusan}
              onChange={handleChange}
              required
              style={styles.input}
            >
              <option value="">Pilih Jurusan</option>
              {jurusanList.map((jur) => (
                <option key={jur.id_jurusan} value={jur.id_jurusan}>
                  {jur.nama_jurusan}
                </option>
              ))}
            </select>
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>Fakultas :</label>
            <select
              name="fakultas"
              value={formData.fakultas}
              onChange={handleChange}
              required
              style={styles.input}
            >
              <option value="">Pilih Fakultas</option>
              {fakultasOptions.map((value) => (
                <option key={value} value={value}>
                  {value}
                </option>
              ))}
            </select>
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>Jenjang :</label>
            <select
              name="jenjang"
              value={formData.jenjang}
              onChange={handleChange}
              required
              style={styles.input}
            >
              <option value="">Pilih Jenjang</option>
              {jenjangOptions.map((value) => (
                <option key={value} value={value}>
                  {value}
                </option>
              ))}
            </select>
          </div>

          <div style={styles.buttonGroup}>
            <button type="button" onClick={handleReset} style={styles.btn}>
              Reset
            </button>
            <button
              type="submit"
              disabled={!isEdit}
              style={{ ...styles.btn, opacity: isEdit ? 1 : 0.5 }}
            >
              Update
            </button>
            <button
              type="submit"
              disabled={isEdit}
              style={{ ...styles.btn, opacity: !isEdit ? 1 : 0.5 }}
            >
              Save
            </button>
            <button
              type="button"
              onClick={handleDelete}
              disabled={!isEdit}
              style={{ ...styles.btn, opacity: isEdit ? 1 : 0.5 }}
            >
              Delete
            </button>
          </div>
        </form>
      </div>

      <div style={styles.rightPanel}>
        <div style={styles.searchContainer}>
          <input
            type="text"
            placeholder="Cari nama / NIM..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            style={styles.searchInput}
          />
        </div>
        <div style={styles.listTitle}>List Data Mahasiswa</div>
        <div style={styles.listBox}>
          {(filteredMahasiswa || []).map((m) => (
            <div
              key={m.id}
              onClick={() => handleItemClick(m)}
              style={{
                ...styles.listItem,
                backgroundColor: selectedId === m.id ? "#e0e0e0" : "#fff",
              }}
            >
              <strong>{m.nama}</strong> ({m.nim}) <br />
              <small>
                {m.jurusan?.jenjang} {m.jurusan?.nama_jurusan} -{" "}
                {m.jurusan?.fakultas}
              </small>
            </div>
          ))}
          {(filteredMahasiswa || []).length === 0 && (
            <p style={{ textValues: "center", color: "#999" }}>
              Data tidak ditemukan
            </p>
          )}
        </div>
        <button
          onClick={handleExportExcel}
          style={{
            marginLeft: "auto",
            marginTop: "10px",
            marginRight: "10px",
            marginBottom: "15px",
            padding: "10px",
            backgroundColor: "#2e7d32",
            color: "white",
            border: "none",
            borderRadius: "5px",
            cursor: "pointer",
          }}
        >
          Export Data ke Excel
        </button>
      </div>
    </div>
  );
}

const styles = {
  container: {
    display: "flex",
    width: "95vw",
    height: "85vh",
    margin: "20px auto",
    border: "2px solid #000",
    fontFamily: "Arial, sans-serif",
  },
  leftPanel: {
    width: "55%",
    borderRight: "2px solid #000",
    display: "flex",
    flexDirection: "column",
  },
  rightPanel: {
    width: "45%",
    padding: "15px",
    display: "flex",
    flexDirection: "column",
  },
  headerForm: {
    borderBottom: "2px solid #000",
    padding: "10px",
    fontSize: "20px",
    fontWeight: "bold",
    textAlign: "center",
  },
  form: {
    padding: "20px",
    flex: 1,
    display: "flex",
    flexDirection: "column",
    justifyContent: "space-between",
  },
  formGroup: { display: "flex", alignItems: "center", marginBottom: "12px" },
  label: { width: "120px", fontWeight: "bold" },
  input: {
    flex: 1,
    padding: "6px",
    border: "1px solid #000",
    fontSize: "14px",
  },
  textarea: {
    flex: 1,
    padding: "6px",
    border: "1px solid #000",
    fontSize: "14px",
    resize: "none",
  },
  buttonGroup: {
    borderTop: "2px solid #000",
    paddingTop: "15px",
    display: "flex",
    justifyContent: "space-around",
  },
  btn: {
    padding: "8px 20px",
    border: "1px solid #000",
    backgroundColor: "#fff",
    cursor: "pointer",
    fontWeight: "bold",
  },
  searchContainer: {
    display: "flex",
    justifyContent: "flex-end",
    marginBottom: "10px",
  },
  searchInput: { width: "60%", padding: "6px", border: "1px solid #000" },
  listTitle: { fontWeight: "bold", marginBottom: "5px" },
  listBox: {
    border: "1px solid #000",
    flex: 1,
    overflowY: "auto",
    padding: "10px",
    backgroundColor: "#f9f9f9",
  },
  listItem: {
    padding: "10px",
    borderBottom: "1px dashed #ccc",
    cursor: "pointer",
    transition: "0.2s",
  },
};

export default App;
