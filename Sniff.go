package main

import (
	"bufio"
	cd "cleandapox"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	lib "request/lib"
	db "xlstosqlite"

	"dapofiles" // "github.com/vibrill/dapofiles"
)

var (
	path                           = "token"
	tolist                         []string
	dataGTK, dataRombel, dataSiswa string
)

const namadb = "sekolah.db"
const header = (`╭━━━╮╱╱╱╱╱╱╱╱╱╱╱╱╭━━━╮╱╱╱╱╱╱╱╭━╮╱╭━╮` + "\n" +
	`╰╮╭╮┃╱╱╱╱╱╱╱╱╱╱╱╱┃╭━╮┃╱╱╱╱╱╱╱┃╭╯╱┃╭╯` + "\n" +
	`╱┃┃┃┃╭━━╮╭━━╮╭━━╮┃╰━━╮╭━╮╱╭╮╭╯╰╮╭╯╰╮` + "\n" +
	`╱┃┃┃┃┃╭╮┃┃╭╮┃┃╭╮┃╰━━╮┃┃╭╮╮┣┫╰╮╭╯╰╮╭╯` + "\n" +
	`╭╯╰╯┃┃╭╮┃┃╰╯┃┃╰╯┃┃╰━╯┃┃┃┃┃┃┃╱┃┃╱╱┃┃` + "\n" +
	`╰━━━╯╰╯╰╯┃╭━╯╰━━╯╰━━━╯╰╯╰╯╰╯╱╰╯╱╱╰╯` + "\n" +
	`╱╱╱╱╱╱╱╱╱┃┃` + "\n" +
	`╱╱╱╱╱╱╱╱╱╰╯` + "\n" +
	`Created by : Vebril OPS Santuy`)

func clear() {
	fmt.Print("\033[H\033[2J")

}

func getdatanpsn() (a string) {
	fmt.Println("Masukkan npsn sekolah Anda: ")
	fmt.Scan(&a)
	clear()
	//fmt.Println(a)
	return a
}

func getdatatoken() (b string) {
	fmt.Println("Masukkan token sekolah Anda: ")
	fmt.Scan(&b)
	clear()
	//fmt.Println(b)
	return b
}

func writext() {
	a := getdatanpsn()
	b := getdatatoken()
	if a != "" && b != "" {
		linesToWrite := a + "\n" + b
		err := ioutil.WriteFile(path, []byte(linesToWrite), 0777)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print("input tidak boleh kosong")
	}

}

func gettoken() (npsn, token string) {

	file, err := os.Open(path)
	if err != nil {
		writext()
		return "", ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tolist = append(tolist, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	npsn = tolist[0]
	token = tolist[1]
	return npsn, token
}

func getdata(perintah string) (text string) {
	npsn, token := gettoken()
	url := "http://localhost:5774/WebService/" + perintah + "?npsn=" + npsn
	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Tidak dapat mengakses link.\n[ERROR] -", err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	text = string([]byte(body))
	//log.Println(text)
	if perintah == "getGtk" {
		dataGTK = text
		lib.JsonPTKtoDB(namadb, dataGTK)
	}
	if perintah == "getRombonganBelajar" {
		dataRombel = text
	}
	if perintah == "getPesertaDidik" {
		dataSiswa = text
		lib.JsonSiswatoDB(namadb, dataSiswa)
	}
	return text

}

func printdata(perintah string) {
	err := ioutil.WriteFile("json\\"+perintah+".json", []byte(getdata(perintah)), 0777)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	clear()
	gettoken() //cek keberadaan file token
	utama()
}

func utama() {
	println("\n" + header)
	println("==============================================\n")
	println("mengakses webserver dapodik...\n")

	db.CreateDB(namadb)
	lib.CreateAllTabble(namadb)
	indikator := [3]string{"getGtk", "getRombonganBelajar", "getPesertaDidik"}
	printdata(indikator[0])
	printdata(indikator[1])
	printdata(indikator[2])
	siswa, guru, tendik := dapofiles.Cek() //cekdapo(downfiles.DownloadFiles())

	println("==============================================\n")
	fmt.Println("ditemukan file dapodik pada folder Download  : ")
	fmt.Println("File siswa didownload pada tanggal : ", siswa[len(siswa)-24:len(siswa)-14])
	fmt.Println("File guru didownload pada tanggal  : ", guru[len(guru)-24:len(guru)-14])
	fmt.Println("File tendik didownload pada tanggal : ", tendik[len(tendik)-24:len(tendik)-14])

	println("==============================================\n")
	cd.Proses()
	db.Proses(namadb)
	lib.CopyFile(namadb)

	println("==============================================\n")
	fmt.Println("Silahkan upload database ", namadb, "pada Skolidbot")
	fmt.Println("Lokasi database berada pada Desktop/DapoSniff/", namadb)
	fmt.Println("Tekan tombol <Enter> untuk mengakhiri proses...")
	fmt.Scanln()
	os.Exit(0)
}
