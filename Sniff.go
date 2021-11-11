package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"dapofiles" // "github.com/vibrill/dapofiles"
)

var (
	path   = "token"
	tolist []string
)

func getdatanpsn() (a string) {
	fmt.Println("Masukkan npsn sekolah Anda: ")
	fmt.Scan(&a)
	//fmt.Println(a)
	return a
}

func getdatatoken() (b string) {
	fmt.Println("Masukkan token sekolah Anda: ")
	fmt.Scan(&b)
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
		fmt.Println("akun telah disimpan")
		fmt.Println("mohon mulai ulang program untuk menjalankan program dengan akun tersebut")
		time.Sleep(3 * time.Second)
		os.Exit(0)
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
	return text

}

func printdata(perintah string) {
	err := ioutil.WriteFile("json\\"+perintah+".json", []byte(getdata(perintah)), 0777)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	println("mohon tunggu, sistem sedang mengakses dapodik")
	indikator := [3]string{"getGtk", "getRombonganBelajar", "getPesertaDidik"}
	printdata(indikator[0])
	printdata(indikator[1])
	printdata(indikator[2])
	println("Data berikut ini telah selesai dibuat dan diletakan dalam folder json:\n1. getGtk.json\n2. getPesertaDidik.json\n3. getRombonganBelajar.json")
	println("silahkan upload tiga file tersebut pada bot Telegram")
	println("memeriksa folder download")
	siswa, guru, tendik := dapofiles.Cek() //cekdapo(downfiles.DownloadFiles())
	fmt.Println("ditemukan file siswa terbaru :\n", siswa)
	fmt.Println("ditemukan file guru terbaru :\n", guru)
	fmt.Println("ditemukan file tendik terbaru :\n", tendik)
	time.Sleep(5 * time.Second)
	os.Exit(0)

}
