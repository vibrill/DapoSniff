package request

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const (
	namaTabelPTK          = "PTKWEB"
	namaTabelRWPangkat    = "RWPANGKATWEB"
	namaTabelRWPendidikan = "RWPENDIDIKANWEB"
)

type Ptk struct {
	Results int    `json:"results"`
	ID      string `json:"id"`
	Start   int    `json:"start"`
	Limit   int    `json:"limit"`
	Rows    []struct {
		Tahun_ajaran_id           string `json:"tahun_ajaran_id"`
		Ptk_terdaftar_id          string `json:"ptk_terdaftar_id"`
		Ptk_id                    string `json:"ptk_id"`
		Ptk_induk                 string `json:"ptk_induk"`
		Tanggal_surat_tugas       string `json:"tanggal_surat_tugas"`
		Nama                      string `json:"nama"`
		Jenis_kelamin             string `json:"jenis_kelamin"`
		Tempat_lahir              string `json:"tempat_lahir"`
		Tanggal_lahir             string `json:"tanggal_lahir"`
		Agama_id                  int    `json:"agama_id"`
		Agama_id_str              string `json:"agama_id_str"`
		Nuptk                     string `json:"nuptk"`
		Nik                       string `json:"nik"`
		Jenis_ptk_id              string `json:"jenis_ptk_id"`
		Jenis_ptk_id_str          string `json:"jenis_ptk_id_str"`
		Status_kepegawaian_id     int    `json:"status_kepegawaian_id"`
		Status_kepegawaian_id_str string `json:"status_kepegawaian_id_str"`
		Nip                       string `json:"nip"`
		Pendidikan_terakhir       string `json:"pendidikan_terakhir"`
		Bidang_studi_terakhir     string `json:"bidang_studi_terakhir"`
		Pangkat_golongan_terakhir string `json:"pangkat_golongan_terakhir"`
		Rwy_pend_formal           []struct {
			Riwayat_pendidikan_formal_id string `json:"riwayat_pendidikan_formal_id"`
			Satuan_pendidikan_formal     string `json:"satuan_pendidikan_formal"`
			Fakultas                     string `json:"fakultas"`
			Kependidikan                 string `json:"kependidikan"`
			Tahun_masuk                  string `json:"tahun_masuk"`
			Tahun_lulus                  string `json:"tahun_lulus"`
			Nim                          string `json:"nim"`
			Status_kuliah                string `json:"status_kuliah"`
			Semester                     string `json:"semester"`
			Ipk                          string `json:"ipk"`
			Prodi                        string `json:"prodi"`
			Id_reg_pd                    string `json:"id_reg_pd"`
			Bidang_studi_id_str          string `json:"bidang_studi_id_str"`
			Jenjang_pendidikan_id_str    string `json:"jenjang_pendidikan_id_str"`
			Gelar_akademik_id_str        string `json:"gelar_akademik_id_str"`
		} `json:"rwy_pend_formal"`
		Rwy_kepangkatan []struct {
			Riwayat_kepangkatan_id  string `json:"riwayat_kepangkatan_id"`
			Nomor_sk                string `json:"nomor_sk"`
			Tanggal_sk              string `json:"tanggal_sk"`
			Tmt_pangkat             string `json:"tmt_pangkat"`
			Masa_kerja_gol_tahun    string `json:"masa_kerja_gol_tahun"`
			Masa_kerja_gol_bulan    string `json:"masa_kerja_gol_bulan"`
			Pangkat_golongan_id_str string `json:"pangkat_golongan_id_str"`
		} `json:"rwy_kepangkatan"`
	} `json:"rows"`
}

var (
	ptkTahunAjaranID           string
	ptkTerdaftarID             string
	ptkID                      string
	ptkInduk                   string
	ptkTanggalSuratTugas       string
	ptkNama                    string
	ptkJenisKelamin            string
	ptkTempatLahir             string
	ptkTanggalLahir            string
	ptkAgamaID                 string
	ptkAgamaIDStr              string
	ptkNuptk                   string
	ptkNik                     string
	ptkJenisPtkID              string
	ptkJenisPtkIDStr           string
	ptkStatusKepegawaianID     string
	ptkStatusKepegawaianIDStr  string
	ptkNip                     string
	ptkPendidikanTerakhir      string
	ptkBidangStudiTerakhir     string
	ptkPangkatGolonganTerakhir string

	ptkRiwayatPendidikanFormalID string
	ptkSatuanPendidikanFormal    string
	ptkFakultas                  string
	ptkKependidikan              string
	ptkTahunMasuk                string
	ptkTahunLulus                string
	ptkNim                       string
	ptkStatusKuliah              string
	ptkSemester                  string
	ptkIpk                       string
	ptkProdi                     string
	ptkIDRegPd                   string
	ptkBidangStudiIDStr          string
	ptkJenjangPendidikanIDStr    string
	ptkGelarAkademikIDStr        string

	ptkRiwayatKepangkatanID string
	ptkNomorSk              string
	ptkTanggalSk            string
	ptkTmtPangkat           string
	ptkMasaKerjaGolTahun    string
	ptkMasaKerjaGolBulan    string
	ptkPangkatGolonganIDStr string
)

func createTableRiwayatKepangkatanPTK(namadb string) {
	db, _ := sql.Open("sqlite3", "./"+namadb) // Op
	defer db.Close()                          // Defer Closing the database
	text := `CREATE TABLE ` + namaTabelRWPangkat + ` (
		"Nama"				TEXT,
		"RiwayatKepangkatanID" TEXT,
		"NomorSk"              TEXT,
		"TanggalSk"            TEXT,
		"TmtPangkat"           TEXT,
		"MasaKerjaGolTahun"    TEXT,
		"MasaKerjaGolBulan"    TEXT,
		"PangkatGolonganIDStr" TEXT
	);`
	//log.Println(text)
	statement, err := db.Prepare(text) // Prepare SQL Statement
	if err != nil {
		log.Fatal("create Table RWPANGKATWEB ERR :" + err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Tabel RWPANGKATWEB telah dibuat")
}

func createTableRiwayatPendidikanPTK(namadb string) {
	db, _ := sql.Open("sqlite3", "./"+namadb) // Op
	defer db.Close()                          // Defer Closing the database
	text := `CREATE TABLE ` + namaTabelRWPendidikan + ` (
		"Nama"				TEXT,
		"RiwayatPendidikanFormalID" TEXT,
		"SatuanPendidikanFormal"    TEXT,
		"Fakultas"                  TEXT,
		"Kependidikan"              TEXT,
		"TahunMasuk"                TEXT,
		"TahunLulus"                TEXT,
		"Nim"                       TEXT,
		"StatusKuliah"              TEXT,
		"Semester"                  TEXT,
		"Ipk"                       TEXT,
		"Prodi"                     TEXT,
		"IDRegPd"                   TEXT,
		"BidangStudiIDStr"          TEXT,
		"JenjangPendidikanIDStr"    TEXT,
		"GelarAkademikIDStr"        TEXT);`

	//log.Println(text)
	statement, err := db.Prepare(text) // Prepare SQL Statement
	if err != nil {
		log.Fatal("create Table RWPENDIDIKANWEB ERR :" + err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Tabel RWPENDIDIKANWEB telah dibuat")
}

func createTablePTK(namadb string) {
	db, _ := sql.Open("sqlite3", "./"+namadb) // Op
	defer db.Close()                          // Defer Closing the database
	text := `CREATE TABLE ` + namaTabelPTK + ` (
		"TahunAjaranID"             TEXT,
		"TerdaftarID"               TEXT,
		"ID"                        TEXT,
		"Induk"                     TEXT,
		"TanggalSuratTugas"         TEXT,
		"Nama"                      TEXT,
		"JenisKelamin"              TEXT,
		"TempatLahir"               TEXT,
		"TanggalLahir"              TEXT,
		"AgamaID"                   TEXT,
		"AgamaIDStr"                TEXT,
		"Nuptk"                     TEXT,
		"Nik"                       TEXT,
		"JenisPtkID"                TEXT,
		"JenisPtkIDStr"             TEXT,
		"StatusKepegawaianID"       TEXT,
		"StatusKepegawaianIDStr"    TEXT,
		"Nip"                       TEXT,
		"PendidikanTerakhir"        TEXT,
		"BidangStudiTerakhir"       TEXT,
		"PangkatGolonganTerakhir"   TEXT);`
	//log.Println(text)
	statement, err := db.Prepare(text) // Prepare SQL Statement
	if err != nil {
		log.Fatal("createTablePTK ERR :" + err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("Tabel PTK telah dibuat")
}

func JsonPTKtoDB(namadb, js string) {
	//init json load
	var jsonString = js
	var jsonData = []byte(jsonString)
	var data Ptk
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//init db access
	db, _ := sql.Open("sqlite3", "./"+namadb) // Open the created SQLite File
	defer db.Close()                          // Defer Closing the database

	countPTK := data.Results
	println("jumlah data PTK ", countPTK)

	for i := 0; i < countPTK; i++ {
		//data setiap row
		//insert data PTK WEB
		ptkTahunAjaranID = formatext(data.Rows[i].Tahun_ajaran_id)
		ptkTerdaftarID = formatext(data.Rows[i].Ptk_terdaftar_id)
		ptkID = formatext(data.Rows[i].Ptk_id)
		ptkInduk = formatext(data.Rows[i].Ptk_induk)
		ptkTanggalSuratTugas = formatext(data.Rows[i].Tanggal_surat_tugas)
		ptkNama = formatext(data.Rows[i].Nama)
		ptkJenisKelamin = formatext(data.Rows[i].Jenis_kelamin)
		ptkTempatLahir = formatext(data.Rows[i].Tempat_lahir)
		ptkTanggalLahir = formatext(data.Rows[i].Tanggal_lahir)
		ptkAgamaID = formatext(strconv.Itoa(data.Rows[i].Agama_id))
		ptkAgamaIDStr = formatext(data.Rows[i].Agama_id_str)
		ptkNuptk = formatext(data.Rows[i].Nuptk)
		ptkNik = formatext(data.Rows[i].Nik)
		ptkJenisPtkID = formatext(data.Rows[i].Jenis_ptk_id)
		ptkJenisPtkIDStr = formatext(data.Rows[i].Jenis_ptk_id_str)
		ptkStatusKepegawaianID = formatext(strconv.Itoa(data.Rows[i].Status_kepegawaian_id))
		ptkStatusKepegawaianIDStr = formatext(data.Rows[i].Status_kepegawaian_id_str)
		ptkNip = formatext(data.Rows[i].Nip)
		ptkPendidikanTerakhir = formatext(data.Rows[i].Pendidikan_terakhir)
		ptkBidangStudiTerakhir = formatext(data.Rows[i].Bidang_studi_terakhir)
		ptkPangkatGolonganTerakhir = formatext(data.Rows[i].Pangkat_golongan_terakhir)

		itemPTK := (ptkTahunAjaranID + ", " +
			ptkTerdaftarID + ", " +
			ptkID + ", " +
			ptkInduk + ", " +
			ptkTanggalSuratTugas + ", " +
			ptkNama + ", " +
			ptkJenisKelamin + ", " +
			ptkTempatLahir + ", " +
			ptkTanggalLahir + ", " +
			ptkAgamaID + ", " +
			ptkAgamaIDStr + ", " +
			ptkNuptk + ", " +
			ptkNik + ", " +
			ptkJenisPtkID + ", " +
			ptkJenisPtkIDStr + ", " +
			ptkStatusKepegawaianID + ", " +
			ptkStatusKepegawaianIDStr + ", " +
			ptkNip + ", " +
			ptkPendidikanTerakhir + ", " +
			ptkBidangStudiTerakhir + ", " +
			ptkPangkatGolonganTerakhir)

		insertSQL := "INSERT INTO " + namaTabelPTK + " VALUES(" + itemPTK + ")"
		statement, err := db.Prepare(insertSQL) // Prepare statement.
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = statement.Exec()
		if err != nil {
			log.Fatalln(err.Error())
		}

		//insert riwayat pendidikan
		countRP := len(data.Rows[i].Rwy_pend_formal)
		for a := 0; a < countRP; a++ {
			ptkRiwayatPendidikanFormalID = formatext(data.Rows[i].Rwy_pend_formal[a].Riwayat_pendidikan_formal_id)
			ptkSatuanPendidikanFormal = formatext(data.Rows[i].Rwy_pend_formal[a].Satuan_pendidikan_formal)
			ptkFakultas = formatext(data.Rows[i].Rwy_pend_formal[a].Fakultas)
			ptkKependidikan = formatext(data.Rows[i].Rwy_pend_formal[a].Kependidikan)
			ptkTahunMasuk = formatext(data.Rows[i].Rwy_pend_formal[a].Tahun_masuk)
			ptkTahunLulus = formatext(data.Rows[i].Rwy_pend_formal[a].Tahun_lulus)
			ptkNim = formatext(data.Rows[i].Rwy_pend_formal[a].Nim)
			ptkStatusKuliah = formatext(data.Rows[i].Rwy_pend_formal[a].Status_kuliah)
			ptkSemester = formatext(data.Rows[i].Rwy_pend_formal[a].Semester)
			ptkIpk = formatext(data.Rows[i].Rwy_pend_formal[a].Ipk)
			ptkProdi = formatext(data.Rows[i].Rwy_pend_formal[a].Prodi)
			ptkIDRegPd = formatext(data.Rows[i].Rwy_pend_formal[a].Id_reg_pd)
			ptkBidangStudiIDStr = formatext(data.Rows[i].Rwy_pend_formal[a].Bidang_studi_id_str)
			ptkJenjangPendidikanIDStr = formatext(data.Rows[i].Rwy_pend_formal[a].Jenjang_pendidikan_id_str)
			ptkGelarAkademikIDStr = formatext(data.Rows[i].Rwy_pend_formal[a].Gelar_akademik_id_str)

			itemRWPEND := (ptkNama + ", " +
				ptkRiwayatPendidikanFormalID + ", " +
				ptkSatuanPendidikanFormal + ", " +
				ptkFakultas + ", " +
				ptkKependidikan + ", " +
				ptkTahunMasuk + ", " +
				ptkTahunLulus + ", " +
				ptkNim + ", " +
				ptkStatusKuliah + ", " +
				ptkSemester + ", " +
				ptkIpk + ", " +
				ptkProdi + ", " +
				ptkIDRegPd + ", " +
				ptkBidangStudiIDStr + ", " +
				ptkJenjangPendidikanIDStr + ", " +
				ptkGelarAkademikIDStr)
			insertSQL := "INSERT INTO " + namaTabelRWPendidikan + " VALUES(" + itemRWPEND + ")"
			statement, err := db.Prepare(insertSQL) // Prepare statement.
			if err != nil {
				log.Fatalln(err.Error())
			}
			_, err = statement.Exec()
			if err != nil {
				log.Fatalln(err.Error())
			}

		}
		countRK := len(data.Rows[i].Rwy_kepangkatan)
		for b := 0; b < countRK; b++ {
			ptkRiwayatKepangkatanID = formatext(data.Rows[i].Rwy_kepangkatan[b].Riwayat_kepangkatan_id)
			ptkNomorSk = formatext(data.Rows[i].Rwy_kepangkatan[b].Nomor_sk)
			ptkTanggalSk = formatext(data.Rows[i].Rwy_kepangkatan[b].Tanggal_sk)
			ptkTmtPangkat = formatext(data.Rows[i].Rwy_kepangkatan[b].Tmt_pangkat)
			ptkMasaKerjaGolTahun = formatext(data.Rows[i].Rwy_kepangkatan[b].Masa_kerja_gol_tahun)
			ptkMasaKerjaGolBulan = formatext(data.Rows[i].Rwy_kepangkatan[b].Masa_kerja_gol_bulan)
			ptkPangkatGolonganIDStr = formatext(data.Rows[i].Rwy_kepangkatan[b].Pangkat_golongan_id_str)

			itemRWPANG := (ptkNama + ", " +
				ptkRiwayatKepangkatanID + ", " +
				ptkNomorSk + ", " +
				ptkTanggalSk + ", " +
				ptkTmtPangkat + ", " +
				ptkMasaKerjaGolTahun + ", " +
				ptkMasaKerjaGolBulan + ", " +
				ptkPangkatGolonganIDStr)

			insertSQL := "INSERT INTO " + namaTabelRWPangkat + " VALUES(" + itemRWPANG + ")"
			statement, err := db.Prepare(insertSQL) // Prepare statement.
			if err != nil {
				log.Fatalln(err.Error())
			}
			_, err = statement.Exec()
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
		println("data ke ", i, "printed")
	}
}
