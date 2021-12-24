package request

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Sekolah struct {
	Results int    `json:"results"`
	ID      string `json:"id"`
	Start   int    `json:"start"`
	Limit   int    `json:"limit"`
	Rows    struct {
		SekolahID   string `json:"sekolah_id"`
		NamaSekolah string `json:"nama"`
		Nss         string `json:"nss"`
		Npsn        string `json:"npsn"`

		BentukPendidikanID    int    `json:"bentuk_pendidikan_id"`
		BentukPendidikanIDStr string `json:"bentuk_pendidikan_id_str"`
		StatusSekolah         string `json:"status_sekolah"`
		StatusSekolahStr      string `json:"status_sekolah_str"`
		AlamatJalan           string `json:"alamat_jalan"`
		Rt                    string `json:"rt"`
		Rw                    string `json:"rw"`
		KodeWilayah           string `json:"kode_wilayah"`
		KodePos               string `json:"kode_pos"`
		NomorTelepon          string `json:"nomor_telepon"`
		NomorFax              string `json:"nomor_fax"`
		Email                 string `json:"email"`
		Website               string `json:"website"`
		IsSks                 bool   `json:"is_sks"`
		Lintang               string `json:"lintang"`
		Bujur                 string `json:"bujur"`
		Dusun                 string `json:"dusun"`
		DesaKelurahan         string `json:"desa_kelurahan"`
		Kecamatan             string `json:"kecamatan"`
		KabupatenKota         string `json:"kabupaten_kota"`
		Provinsi              string `json:"provinsi"`
	} `json:"rows"`
}

var (
	SekolahId             string
	NamaSekolah           string
	Nss                   string
	Npsn                  string
	BentukPendidikanID    string
	BentukPendidikanIDStr string
	StatusSekolah         string
	StatusSekolahStr      string
	AlamatJalan           string
	Rt                    string
	Rw                    string
	KodeWilayah           string
	KodePos               string
	NomorTelepon          string
	NomorFax              string
	Email                 string
	Website               string
	IsSks                 string
	Lintang               string
	Bujur                 string
	Dusun                 string
	DesaKelurahan         string
	Kecamatan             string
	KabupatenKota         string
	Provinsi              string
)

func createTableSekolah(db *sql.DB) {
	text := `CREATE TABLE SEKOLAH (
		"SekolahId"            TEXT,
		"NamaSekolah"           TEXT,
		"Nss"                   TEXT,
		"Npsn"                  TEXT,
		"BentukPendidikanId"    TEXT,
		"BentukPendidikanIdStr" TEXT,
		"StatusSekolah"         TEXT,
		"StatusSekolahStr"      TEXT,
		"AlamatJalan"           TEXT,
		"Rt"                    TEXT,
		"Rw"                    TEXT,
		"KodeWilayah"           TEXT,
		"KodePos"               TEXT,
		"NomorTelepon"          TEXT,
		"NomorFax"              TEXT,
		"Email"                 TEXT,
		"Website"               TEXT,
		"IsSks"                 TEXT,
		"Lintang"               TEXT,
		"Bujur"                 TEXT,
		"Dusun"                 TEXT,
		"DesaKelurahan"         TEXT,
		"Kecamatan"             TEXT,
		"KabupatenKota"         TEXT,
		"Provinsi"              TEXT);`

	statement, err := db.Prepare(text) // Prepare SQL Statement
	if err != nil {
		log.Fatal("create db sekolah" + err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func JsonSekolahtoDB(db *sql.DB, js string) {
	//init json load
	var jsonString = js
	var jsonData = []byte(jsonString)
	var data Sekolah
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var text string

	SekolahId = formatext(data.Rows.SekolahID)
	NamaSekolah = formatext(data.Rows.NamaSekolah)
	Nss = formatext(data.Rows.Nss)
	Npsn = formatext(data.Rows.Npsn)
	BentukPendidikanID = formatext(strconv.Itoa(data.Rows.BentukPendidikanID))
	BentukPendidikanIDStr = formatext(data.Rows.BentukPendidikanIDStr)
	StatusSekolah = formatext(data.Rows.StatusSekolah)
	StatusSekolahStr = formatext(data.Rows.StatusSekolahStr)
	AlamatJalan = formatext(data.Rows.AlamatJalan)
	Rt = formatext(data.Rows.Rt)
	Rw = formatext(data.Rows.Rw)
	KodeWilayah = formatext(data.Rows.KodeWilayah)
	KodePos = formatext(data.Rows.KodePos)
	NomorTelepon = formatext(data.Rows.NomorTelepon)
	NomorFax = formatext(data.Rows.NomorFax)
	Email = formatext(data.Rows.Email)
	Website = formatext(data.Rows.Website)
	IsSks = formatext(strconv.FormatBool(data.Rows.IsSks))
	Lintang = formatext(data.Rows.Lintang)
	Bujur = formatext(data.Rows.Bujur)
	Dusun = formatext(data.Rows.Dusun)
	DesaKelurahan = formatext(data.Rows.DesaKelurahan)
	Kecamatan = formatext(data.Rows.Kecamatan)
	KabupatenKota = formatext(data.Rows.KabupatenKota)
	Provinsi = formatext(data.Rows.Provinsi)

	text = (SekolahId + ", " +
		NamaSekolah + ", " +
		Nss + ", " +
		Npsn + ", " +
		BentukPendidikanID + ", " +
		BentukPendidikanIDStr + ", " +
		StatusSekolah + ", " +
		StatusSekolahStr + ", " +
		AlamatJalan + ", " +
		Rt + ", " +
		Rw + ", " +
		KodeWilayah + ", " +
		KodePos + ", " +
		NomorTelepon + ", " +
		NomorFax + ", " +
		Email + ", " +
		Website + ", " +
		IsSks + ", " +
		Lintang + ", " +
		Bujur + ", " +
		Dusun + ", " +
		DesaKelurahan + ", " +
		Kecamatan + ", " +
		KabupatenKota + ", " +
		Provinsi)

	insertSQL := "INSERT INTO SEKOLAH VALUES(" + text + ")"
	statement, err := db.Prepare(insertSQL) // Prepare statement.
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
