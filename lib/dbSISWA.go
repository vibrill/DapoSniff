package request

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Siswa struct {
	Results int    `json:"results"`
	ID      string `json:"id"`
	Start   int    `json:"start"`
	Limit   int    `json:"limit"`
	Rows    []struct {
		RegistrasiID          string `json:"registrasi_id"`
		JenisPendaftaranID    string `json:"jenis_pendaftaran_id"`
		JenisPendaftaranIDStr string `json:"jenis_pendaftaran_id_str"`
		Nipd                  string `json:"nipd"`
		TanggalMasukSekolah   string `json:"tanggal_masuk_sekolah"`
		SekolahAsal           string `json:"sekolah_asal"`
		PesertaDidikID        string `json:"peserta_didik_id"`
		Nama                  string `json:"nama"`
		Nisn                  string `json:"nisn"`
		JenisKelamin          string `json:"jenis_kelamin"`
		Nik                   string `json:"nik"`
		TempatLahir           string `json:"tempat_lahir"`
		TanggalLahir          string `json:"tanggal_lahir"`
		AgamaID               int    `json:"agama_id"` //int
		AgamaIDStr            string `json:"agama_id_str"`
		AlamatJalan           string `json:"alamat_jalan"`
		NomorTeleponRumah     string `json:"nomor_telepon_rumah"`
		NomorTeleponSeluler   string `json:"nomor_telepon_seluler"`
		NamaAyah              string `json:"nama_ayah"`
		PekerjaanAyahID       int    `json:"pekerjaan_ayah_id"` //int
		PekerjaanAyahIDStr    string `json:"pekerjaan_ayah_id_str"`
		NamaIbu               string `json:"nama_ibu"`
		PekerjaanIbuID        int    `json:"pekerjaan_ibu_id"` //int
		PekerjaanIbuIDStr     string `json:"pekerjaan_ibu_id_str"`
		NamaWali              string `json:"nama_wali"`
		PekerjaanWaliID       int    `json:"pekerjaan_wali_id"` //int
		PekerjaanWaliIDStr    string `json:"pekerjaan_wali_id_str"`
		SemesterID            string `json:"semester_id"`
		Email                 string `json:"email"`
		AnggotaRombelID       string `json:"anggota_rombel_id"`
		RombonganBelajarID    string `json:"rombongan_belajar_id"`
		TingkatPendidikanID   string `json:"tingkat_pendidikan_id"`
		NamaRombel            string `json:"nama_rombel"`
		KurikulumID           int    `json:"kurikulum_id"` //int
		KurikulumIDStr        string `json:"kurikulum_id_str"`
	} `json:"rows"`
}

var (
	PDRegistrasiID          string
	PDJenisPendaftaranID    string
	PDJenisPendaftaranIDStr string
	PDNipd                  string
	PDTanggalMasukSekolah   string
	PDSekolahAsal           string
	PDPesertaDidikID        string
	PDNama                  string
	PDNisn                  string
	PDJenisKelamin          string
	PDNik                   string
	PDTempatLahir           string
	PDTanggalLahir          string
	PDAgamaID               string //int
	PDAgamaIDStr            string
	PDAlamatJalan           string
	PDNomorTeleponRumah     string
	PDNomorTeleponSeluler   string
	PDNamaAyah              string
	PDPekerjaanAyahID       string //int
	PDPekerjaanAyahIDStr    string
	PDNamaIbu               string
	PDPekerjaanIbuID        string //int
	PDPekerjaanIbuIDStr     string
	PDNamaWali              string
	PDPekerjaanWaliID       string //int
	PDPekerjaanWaliIDStr    string
	PDSemesterID            string
	PDEmail                 string
	PDAnggotaRombelID       string
	PDRombonganBelajarID    string
	PDTingkatPendidikanID   string
	PDNamaRombel            string
	PDKurikulumID           string //int
	PDKurikulumIDStr        string
)

func createTableSiswa(namadb string) {
	db, _ := sql.Open("sqlite3", "./"+namadb) // Op
	defer db.Close()                          // Defer Closing the database
	text := `CREATE TABLE SISWAWEB (
		"RegistrasiID"          TEXT,
		"JenisPendaftaranID"    TEXT,
		"JenisPendaftaranIDStr" TEXT,
		"Nipd"                  TEXT,
		"TanggalMasukSekolah"   TEXT,
		"SekolahAsal"           TEXT,
		"PesertaDidikID"        TEXT,
		"Nama"                  TEXT,
		"Nisn"                  TEXT,
		"JenisKelamin"          TEXT,
		"Nik"                   TEXT,
		"TempatLahir"           TEXT,
		"TanggalLahir"          TEXT,
		"AgamaID"               TEXT,
		"AgamaIDStr"            TEXT,
		"AlamatJalan"           TEXT,
		"NomorTeleponRumah"     TEXT,
		"NomorTeleponSeluler"   TEXT,
		"NamaAyah"              TEXT,
		"PekerjaanAyahID"       TEXT,
		"PekerjaanAyahIDStr"    TEXT,
		"NamaIbu"               TEXT,
		"PekerjaanIbuID"        TEXT,
		"PekerjaanIbuIDStr"     TEXT,
		"NamaWali"              TEXT,
		"PekerjaanWaliID"       TEXT,
		"PekerjaanWaliIDStr"    TEXT,
		"SemesterID"            TEXT,
		"Email"                 TEXT,
		"AnggotaRombelID"       TEXT,
		"RombonganBelajarID"    TEXT,
		"TingkatPendidikanID"   TEXT,
		"NamaRombel"            TEXT,
		"KurikulumID"           TEXT,
		"KurikulumIDStr"        TEXT);`

	//log.Println(text)
	statement, err := db.Prepare(text) // Prepare SQL Statement
	if err != nil {
		log.Fatal("keempat" + err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Tabel siswa telah dibuat")
}

func formatext(a string) (b string) {
	b = `"` + a + `"`
	return b
}

func JsonSiswatoDB(namadb, js string) {
	//init json load
	var jsonString = js
	var jsonData = []byte(jsonString)
	var data Siswa
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var text string

	//init db access
	db, _ := sql.Open("sqlite3", "./"+namadb) // Open the created SQLite File
	defer db.Close()                          // Defer Closing the database

	count := data.Results
	for i := 0; i < count; i++ { //data setiap row
		PDRegistrasiID = formatext(data.Rows[i].RegistrasiID)
		PDJenisPendaftaranID = formatext(data.Rows[i].JenisPendaftaranID)
		PDJenisPendaftaranIDStr = formatext(data.Rows[i].JenisPendaftaranIDStr)
		PDNipd = formatext(data.Rows[i].Nipd)
		PDTanggalMasukSekolah = formatext(data.Rows[i].TanggalMasukSekolah)
		PDSekolahAsal = formatext(data.Rows[i].SekolahAsal)
		PDPesertaDidikID = formatext(data.Rows[i].PesertaDidikID)
		PDNama = formatext(data.Rows[i].Nama)
		PDNisn = formatext(data.Rows[i].Nisn)
		PDJenisKelamin = formatext(data.Rows[i].JenisKelamin)
		PDNik = formatext(data.Rows[i].Nik)
		PDTempatLahir = formatext(data.Rows[i].TempatLahir)
		PDTanggalLahir = formatext(data.Rows[i].TanggalLahir)
		PDAgamaID = formatext(strconv.Itoa(data.Rows[i].AgamaID))
		PDAgamaIDStr = formatext(data.Rows[i].AgamaIDStr)
		PDAlamatJalan = formatext(data.Rows[i].AlamatJalan)
		PDNomorTeleponRumah = formatext(data.Rows[i].NomorTeleponRumah)
		PDNomorTeleponSeluler = formatext(data.Rows[i].NomorTeleponSeluler)
		PDNamaAyah = formatext(data.Rows[i].NamaAyah)
		PDPekerjaanAyahID = formatext(strconv.Itoa(data.Rows[i].PekerjaanAyahID))
		PDPekerjaanAyahIDStr = formatext(data.Rows[i].PekerjaanAyahIDStr)
		PDNamaIbu = formatext(data.Rows[i].NamaIbu)
		PDPekerjaanIbuID = formatext(strconv.Itoa(data.Rows[i].PekerjaanIbuID))
		PDPekerjaanIbuIDStr = formatext(data.Rows[i].PekerjaanIbuIDStr)
		PDNamaWali = formatext(data.Rows[i].NamaWali)
		PDPekerjaanWaliID = formatext(strconv.Itoa(data.Rows[i].PekerjaanWaliID))
		PDPekerjaanWaliIDStr = formatext(data.Rows[i].PekerjaanWaliIDStr)
		PDSemesterID = formatext(data.Rows[i].SemesterID)
		PDEmail = formatext(data.Rows[i].Email)
		PDAnggotaRombelID = formatext(data.Rows[i].AnggotaRombelID)
		PDRombonganBelajarID = formatext(data.Rows[i].RombonganBelajarID)
		PDTingkatPendidikanID = formatext(data.Rows[i].TingkatPendidikanID)
		PDNamaRombel = formatext(data.Rows[i].NamaRombel)
		PDKurikulumID = formatext(strconv.Itoa(data.Rows[i].KurikulumID))
		PDKurikulumIDStr = formatext(data.Rows[i].KurikulumIDStr)

		text = (PDRegistrasiID + ", " +
			PDJenisPendaftaranID + ", " +
			PDJenisPendaftaranIDStr + ", " +
			PDNipd + ", " +
			PDTanggalMasukSekolah + ", " +
			PDSekolahAsal + ", " +
			PDPesertaDidikID + ", " +
			PDNama + ", " +
			PDNisn + ", " +
			PDJenisKelamin + ", " +
			PDNik + ", " +
			PDTempatLahir + ", " +
			PDTanggalLahir + ", " +
			PDAgamaID + ", " +
			PDAgamaIDStr + ", " +
			PDAlamatJalan + ", " +
			PDNomorTeleponRumah + ", " +
			PDNomorTeleponSeluler + ", " +
			PDNamaAyah + ", " +
			PDPekerjaanAyahID + ", " +
			PDPekerjaanAyahIDStr + ", " +
			PDNamaIbu + ", " +
			PDPekerjaanIbuID + ", " +
			PDPekerjaanIbuIDStr + ", " +
			PDNamaWali + ", " +
			PDPekerjaanWaliID + ", " +
			PDPekerjaanWaliIDStr + ", " +
			PDSemesterID + ", " +
			PDEmail + ", " +
			PDAnggotaRombelID + ", " +
			PDRombonganBelajarID + ", " +
			PDTingkatPendidikanID + ", " +
			PDNamaRombel + ", " +
			PDKurikulumID + ", " +
			PDKurikulumIDStr)
		insertSQL := "INSERT INTO SISWAWEB VALUES(" + text + ")"
		statement, err := db.Prepare(insertSQL) // Prepare statement.
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = statement.Exec()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	println("one")
}
