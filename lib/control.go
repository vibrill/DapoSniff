package request

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateAllTabble(db *sql.DB) {
	createTableSekolah(db)
	createTablePTK(db)
	createTableRiwayatKepangkatanPTK(db)
	createTableRiwayatPendidikanPTK(db)
	createTableSiswa(db)
	createTabelRombel(db)
	createTabelAnggotaRombel(db)
	createTablePembelajaran(db)
}
