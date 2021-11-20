package request

func CreateAllTabble(namadb string) {
	createTablePTK(namadb)
	createTableRiwayatKepangkatanPTK(namadb)
	createTableRiwayatPendidikanPTK(namadb)
	createTableSiswa(namadb)
}
