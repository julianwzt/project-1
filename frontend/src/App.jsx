import { useState, useEffect } from "react";
import axios from "axios";

const API_URL = "http://localhost:8080/api/mahasiswa";

function App() {
  const [mahasiswa, setMahasiswa] = useState([]);
  const [searchQuery, setSearchQuery] = useState("");
  const [isEdit, setIsEdit] = useState(false);
  const [selectedId, setSelectedId] = useState(null);

  const [form, setForm] = useState({
    nama: "",
    umur: "",
    nim: "",
    tgl_lahir: "",
    alamat: "",
    id_jurusan: "1",
  });

  const fetchMahasiswa = async () => {
    try {
      const res = await axios.get(API_URL);
      setMahasiswa(res.data || []);
    } catch (err) {
      console.error("Gagal mengambil data:", err);
    }
  };

  useEffect(() => {
    fetchMahasiswa();
  }, []);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleReset = () => {
    setForm({
      nama: "",
      umur: "",
      nim: "",
      tgl_lahir: "",
      alamat: "",
      id_jurusan: "1",
    });
    setIsEdit(false);
    setSelectedId(null);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = {
      ...form,
      umur: parseInt(form.umur),
      id_jurusan: parseInt(form.id_jurusan),
      tgl_lahir: form.tgl_lahir ? new Date(form.tgl_lahir).toISOString() : "",
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
      alert("Gagal memproses data: " + err.response?.data?.error);
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
    setForm({
      nama: m.nama,
      umur: m.umur,
      nim: m.nim,
      tgl_lahir: dateOnly,
      alamat: m.alamat,
      id_jurusan: String(m.id_jurusan),
    });
  };

  const filteredMahasiswa = mahasiswa.filter(
    (m) =>
      m.nama.toLowerCase().includes(searchQuery.toLowerCase()) ||
      m.nim.includes(searchQuery),
  );

  return (
    <div style={styles.container}>
      <div style={styles.leftPanel}>
        <div style={styles.headerForm}>Judul Form</div>
        <form onSubmit={handleSubmit} style={styles.form}>
          <div style={styles.formGroup}>
            <label style={styles.label}>Nama :</label>
            <input
              type="text"
              name="nama"
              value={form.nama}
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
              value={form.umur}
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
              value={form.nim}
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
              value={form.tgl_lahir}
              onChange={handleChange}
              required
              style={styles.input}
            />
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>Alamat :</label>
            <textarea
              name="alamat"
              value={form.alamat}
              onChange={handleChange}
              rows="3"
              style={styles.textarea}
            />
          </div>
          <div style={styles.formGroup}>
            <label style={styles.label}>Jurusan :</label>
            <select
              name="id_jurusan"
              value={form.id_jurusan}
              onChange={handleChange}
              style={styles.input}
            >
              <option value="1">Teknik Informatika</option>
              <option value="2">Sistem Informasi</option>
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
            placeholder="🔍 Cari nama / NIM..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            style={styles.searchInput}
          />
        </div>
        <div style={styles.listTitle}>List Data Mahasiswa</div>
        <div style={styles.listBox}>
          {filteredMahasiswa.map((m) => (
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
          {filteredMahasiswa.length === 0 && (
            <p style={{ textValues: "center", color: "#999" }}>
              Data tidak ditemukan
            </p>
          )}
        </div>
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
