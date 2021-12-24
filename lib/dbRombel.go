package request

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const (
	namaTabelRombel        = "ROMBELWEB"
	namaTabelAnggotaRombel = "ANGGOTAROMBELWEB"
	namaTabelPembelajaran  = "PEMBELAJARANWEB"
)

type Rombel struct {
	Results int    `json:"results"`
	ID      string `json:"id"`
	Start   int    `json:"start"`
	Limit   int    `json:"limit"`
	Rows    []struct {
		RombonganBelajarID  string `json:"rombongan_belajar_id"`
		Nama                string `json:"nama"`
		TingkatPendidikanID string `json:"tingkat_pendidikan_id"`
		SemesterID          string `json:"semester_id"`
		JenisRombel         string `json:"jenis_rombel"`
		KurikulumID         int    `json:"kurikulum_id"`
		KurikulumIDStr      string `json:"kurikulum_id_str"`
		IDRuang             string `json:"id_ruang"`
		IDRuangStr          string `json:"id_ruang_str"`
		MovingClass         string `json:"moving_class"`
		RombelPtkID         string `json:"ptk_id"`
		PtkIDStr            string `json:"ptk_id_str"`
		JenisRombelStr      string `json:"jenis_rombel_str"`
		AnggotaRombel       []struct {
			AnggotaRombelID       string `json:"anggota_rombel_id"`
			PesertaDidikID        string `json:"peserta_didik_id"`
			JenisPendaftaranID    string `json:"jenis_pendaftaran_id"`
			JenisPendaftaranIDStr string `json:"jenis_pendaftaran_id_str"`
		} `json:"anggota_rombel"`
		Pembelajaran []struct {
			PembelajaranID       string `json:"pembelajaran_id"`
			MataPelajaranID      int    `json:"mata_pelajaran_id"`
			MataPelajaranIDStr   string `json:"mata_pelajaran_id_str"`
			PtkTerdaftarID       string `json:"ptk_terdaftar_id"`
			PembelajaranPtkID    string `json:"ptk_id"`
			NamaMataPelajaran    string `json:"nama_mata_pelajaran"`
			IndukPembelajaranID  string `json:"induk_pembelajaran_id"`
			JamMengajarPerMinggu string `json:"jam_mengajar_per_minggu"`
			StatusDiKurikulum    string `json:"status_di_kurikulum"`
			StatusDiKurikulumStr string `json:"status_di_kurikulum_str"`
		} `json:"pembelajaran"`
	} `json:"rows"`
}

var (
	Results             int
	ID                  string
	Start               int
	Limit               int
	RombonganBelajarID  string
	Nama                string
	TingkatPendidikanID string
	SemesterID          string
	JenisRombel         string
	KurikulumID         string
	KurikulumIDStr      string
	IDRuang             string
	IDRuangStr          string
	MovingClass         string
	RombelPtkID         string
	PtkIDStr            string
	JenisRombelStr      string

	AnggotaRombelID       string
	PesertaDidikID        string
	JenisPendaftaranID    string
	JenisPendaftaranIDStr string

	PembelajaranID       string
	MataPelajaranID      string
	MataPelajaranIDStr   string
	PtkTerdaftarID       string
	PembelajaranPtkID    string
	NamaMataPelajaran    string
	IndukPembelajaranID  string
	JamMengajarPerMinggu string
	StatusDiKurikulum    string
	StatusDiKurikulumStr string
)

func createTabelRombel(db *sql.DB) {
	/* db, _ := sql.Open("sqlite3", "./"+namadb) // Op
	defer db.Close()                          // Defer Closing the database */
	text := `CREATE TABLE ` + namaTabelRombel + ` (
		"RombonganBelajarID"  TEXT, 
		"Nama"                TEXT, 
		"TingkatPendidikanID" TEXT,
		"SemesterID"          TEXT, 
		"JenisRombel"         TEXT, 
		"KurikulumID"         TEXT,    
		"KurikulumIDStr"      TEXT, 
		"IDRuang"             TEXT, 
		"IDRuangStr"          TEXT, 
		"MovingClass"         TEXT,
		"RombelPtkID"               TEXT, 
		"PtkIDStr"            TEXT, 
		"JenisRombelStr"      TEXT);`

	//fmt.Println(text)
	statement, err := db.Prepare(text) // Prepare SQL Statement
	if err != nil {
		log.Fatal("create Table " + namaTabelRombel + " ERR :" + err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func createTabelAnggotaRombel(db *sql.DB) {
	/* db, _ := sql.Open("sqlite3", "./"+namadb) // Op
	defer db.Close()                          // Defer Closing the database */
	text := `CREATE TABLE ` + namaTabelAnggotaRombel + ` (
		"Nama"                TEXT,
		"PtkIDStr"            TEXT,
		"AnggotaRombelID"       TEXT, 
		"PesertaDidikID"        TEXT, 
		"JenisPendaftaranID"    TEXT, 
		"JenisPendaftaranIDStr" TEXT);`

	//fmt.Println(text)
	statement, err := db.Prepare(text) // Prepare SQL Statement
	if err != nil {
		log.Fatal("create Table " + namaTabelAnggotaRombel + " ERR :" + err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func createTablePembelajaran(db *sql.DB) {
	/* db, _ := sql.Open("sqlite3", "./"+namadb) // Op
	defer db.Close()                          // Defer Closing the database */
	text := `CREATE TABLE ` + namaTabelPembelajaran + ` (
		"Nama"                TEXT,
		"PtkIDStr"            TEXT,
		"PembelajaranID"       TEXT, 
		"MataPelajaranID"      TEXT, 
		"MataPelajaranIDStr"   TEXT,
		"PtkTerdaftarID"       TEXT, 
		"PembelajaranPtkID"                TEXT, 
		"NamaMataPelajaran"    TEXT, 
		"IndukPembelajaranID"  TEXT, 
		"JamMengajarPerMinggu" TEXT, 
		"StatusDiKurikulum"    TEXT, 
		"StatusDiKurikulumStr" TEXT );`
	//fmt.Println(text)
	statement, err := db.Prepare(text) // Prepare SQL Statement
	if err != nil {
		log.Fatal("create  " + namaTabelPembelajaran + " ERR :" + err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func JsonRombelToDB(db *sql.DB, js string) {
	//init json load
	var jsonString = js
	var jsonData = []byte(jsonString)
	var data Rombel
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalln(err.Error())
	}

	countRombel := data.Results
	//println(countPTK, "PTK ditemukan")

	for i := 0; i < countRombel; i++ {
		//data setiap row
		//insert data PTK WEB
		RombonganBelajarID = formatext(data.Rows[i].RombonganBelajarID)
		Nama = formatext(data.Rows[i].Nama)
		TingkatPendidikanID = formatext(data.Rows[i].TingkatPendidikanID)
		SemesterID = formatext(data.Rows[i].SemesterID)
		JenisRombel = formatext(data.Rows[i].JenisRombel)
		KurikulumID = formatext(strconv.Itoa(data.Rows[i].KurikulumID))
		KurikulumIDStr = formatext(data.Rows[i].KurikulumIDStr)
		IDRuang = formatext(data.Rows[i].IDRuang)
		IDRuangStr = formatext(data.Rows[i].IDRuangStr)
		MovingClass = formatext(data.Rows[i].MovingClass)
		RombelPtkID = formatext(data.Rows[i].RombelPtkID)
		PtkIDStr = formatext(data.Rows[i].PtkIDStr)
		JenisRombelStr = formatext(data.Rows[i].JenisRombelStr)

		itemRombel := (RombonganBelajarID + ", " +
			Nama + ", " +
			TingkatPendidikanID + ", " +
			SemesterID + ", " +
			JenisRombel + ", " +
			KurikulumID + ", " +
			KurikulumIDStr + ", " +
			IDRuang + ", " +
			IDRuangStr + ", " +
			MovingClass + ", " +
			RombelPtkID + ", " +
			PtkIDStr + ", " +
			JenisRombelStr)

		insertSQL := "INSERT INTO " + namaTabelRombel + " VALUES(" + itemRombel + ")"
		statement, err := db.Prepare(insertSQL) // Prepare statement.
		if err != nil {
			log.Fatalln(err.Error())
		}
		_, err = statement.Exec()
		if err != nil {
			log.Fatalln(err.Error())
		}

		//insert riwayat pendidikan
		countAR := len(data.Rows[i].AnggotaRombel)
		for a := 0; a < countAR; a++ {
			Nama = formatext(data.Rows[i].Nama)
			PtkIDStr = formatext(data.Rows[i].PtkIDStr)
			AnggotaRombelID = formatext(data.Rows[i].AnggotaRombel[a].AnggotaRombelID)
			PesertaDidikID = formatext(data.Rows[i].AnggotaRombel[a].PesertaDidikID)
			JenisPendaftaranID = formatext(data.Rows[i].AnggotaRombel[a].JenisPendaftaranID)
			JenisPendaftaranIDStr = formatext(data.Rows[i].AnggotaRombel[a].JenisPendaftaranIDStr)

			itemAR := (Nama + ", " +
				PtkIDStr + ", " +
				AnggotaRombelID + ", " +
				PesertaDidikID + ", " +
				JenisPendaftaranID + ", " +
				JenisPendaftaranIDStr)
			insertSQL := "INSERT INTO " + namaTabelAnggotaRombel + " VALUES(" + itemAR + ")"
			statement, err := db.Prepare(insertSQL) // Prepare statement.
			if err != nil {
				log.Fatalln(err.Error())
			}
			_, err = statement.Exec()
			if err != nil {
				log.Fatalln(err.Error())
			}

		}
		countPemb := len(data.Rows[i].Pembelajaran)
		for b := 0; b < countPemb; b++ {
			Nama = formatext(data.Rows[i].Nama)
			PtkIDStr = formatext(data.Rows[i].PtkIDStr)
			PembelajaranID = formatext(data.Rows[i].Pembelajaran[b].PembelajaranID)
			MataPelajaranID = formatext(strconv.Itoa(data.Rows[i].Pembelajaran[b].MataPelajaranID))
			MataPelajaranIDStr = formatext(data.Rows[i].Pembelajaran[b].MataPelajaranIDStr)
			PtkTerdaftarID = formatext(data.Rows[i].Pembelajaran[b].PtkTerdaftarID)
			PembelajaranPtkID = formatext(data.Rows[i].Pembelajaran[b].PembelajaranPtkID)
			NamaMataPelajaran = formatext(data.Rows[i].Pembelajaran[b].NamaMataPelajaran)
			IndukPembelajaranID = formatext(data.Rows[i].Pembelajaran[b].IndukPembelajaranID)
			JamMengajarPerMinggu = formatext(data.Rows[i].Pembelajaran[b].JamMengajarPerMinggu)
			StatusDiKurikulum = formatext(data.Rows[i].Pembelajaran[b].StatusDiKurikulum)
			StatusDiKurikulumStr = formatext(data.Rows[i].Pembelajaran[b].StatusDiKurikulumStr)

			itemPemb := (Nama + ", " +
				PtkIDStr + ", " +
				PembelajaranID + ", " +
				MataPelajaranID + ", " +
				MataPelajaranIDStr + ", " +
				PtkTerdaftarID + ", " +
				PembelajaranPtkID + ", " +
				NamaMataPelajaran + ", " +
				IndukPembelajaranID + ", " +
				JamMengajarPerMinggu + ", " +
				StatusDiKurikulum + ", " +
				StatusDiKurikulumStr)

			insertSQL := "INSERT INTO " + namaTabelPembelajaran + " VALUES(" + itemPemb + ")"
			statement, err := db.Prepare(insertSQL) // Prepare statement.
			if err != nil {
				log.Fatalln(err.Error())
			}
			_, err = statement.Exec()
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
	}
}
