package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func CreateAllTabble(namadb string) {
	createTablePTK(namadb)
	createTableRiwayatKepangkatanPTK(namadb)
	createTableRiwayatPendidikanPTK(namadb)
	createTableSiswa(namadb)
}

func CopyFile(namadb string) { //old  dan new harus path file
	u, _ := os.UserHomeDir()
	_, err := ioutil.ReadDir(string(u) + `/Desktop/DapoSniff`)
	path := u + `/Desktop/DapoSniff`
	if err != nil {
		_, err = ioutil.ReadDir(`E:/Desktop/DapoSniff`)
		path = `E:/Desktop/DapoSniff`
		if err != nil {
			_, err = ioutil.ReadDir(`D:/Desktop/DapoSniff`)
			path = `D:/Desktop/DapoSniff`
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	original, err := os.Open(namadb)
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	//make file
	new, err := os.Create(path + `/` + namadb)
	if err != nil {
		log.Fatal(err)
	}
	defer new.Close()

	//copy file
	_, err = io.Copy(new, original) // _ = bytesWritten
	if err != nil {
		log.Fatal(err)
	}
}
