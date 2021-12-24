package main

import (
	"bufio"
	cd "cleandapox"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	lib "request/lib"
	xs "xlstosqlite"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/browser"
	// "github.com/vibrill/dapofiles"
)

var (
	keypath                                     = "key"
	tolist                                      []string
	dataGTK, dataRombel, dataSiswa, dataSekolah string
	link                                        string = "https://filmier-wren-1640.dataplicity.io/"
	scramble                                    string = ".W3.{aC.9@n.Pr7.o-.le.W.F8.a.)..{Ct./._.te..)VLNd.@.35. k33C.sbb..<9.y....UW gPSE.d:u.yFk. 3.uR-bV..chBN-ld.B|.B.@..}8ri2k+z.i.v..C<l.=#.q,l.9#c.8ia.Hn.Vars^..!n}.z.{.t_..(WM AP$!2.O/@l.|YP....S!SFHEURW .qE_DkS..fd..a.6j.&Hn97..5.tF.Nu.$u6(...7y"
	namadb                                      string = "file"
	namads                                      string = "data.sb"
	scrapas                                     string = "...Sjs.k.t.u.'ll.kgl.kyB.kyo.iyt.ky2.ki0.hd2.jr0..."
	isarsip                                     string
)

const (
	identitas = `Created by : V-Brilliant OPS Santuy`
)

func getdatanpsn() (a string) {
	fmt.Println("Masukkan npsn sekolah Anda: ")
	fmt.Scan(&a)
	return a
}

func getdatatoken() (b string) {
	fmt.Println("Masukkan token akses dapodik Anda: ")
	fmt.Scan(&b)
	return b
}

func getarsip() (c string) {
	fmt.Println("arsip database Y/N (default = N): ")
	fmt.Scan(&c)
	if c == "Y" || c == "y" {
		c = "Y"
	} else {
		c = "N"
	}
	return c
}

func writext() {
	a := getdatanpsn()
	b := getdatatoken()
	c := getarsip()
	if a != "" && b != "" {
		linesToWrite := a + "\n" + b + "\n" + c
		err := ioutil.WriteFile(keypath, []byte(linesToWrite), 0777)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print("input tidak boleh kosong")
	}

}

func gettoken() (npsn, token, arsip string) {

	file, err := os.Open(keypath)
	if err != nil {
		writext()
		return "", "", ""
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
	arsip = tolist[2]
	return npsn, token, arsip
}

func getdata(db *sql.DB, perintah string) (text string) {
	npsn, token, arsip := gettoken()
	isarsip = arsip
	url := "http://localhost:5774/WebService/" + perintah + "?npsn=" + npsn
	//println(url)
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
		lib.JsonPTKtoDB(db, dataGTK)
	}
	if perintah == "getRombonganBelajar" {
		dataRombel = text
		lib.JsonRombelToDB(db, dataRombel)
	}
	if perintah == "getPesertaDidik" {
		dataSiswa = text
		lib.JsonSiswatoDB(db, dataSiswa)
	}
	if perintah == "getSekolah" {
		dataSekolah = text
		lib.JsonSekolahtoDB(db, dataSekolah)
	}
	return text

}

func printdata(db *sql.DB, perintah string) {
	getdata(db, perintah)
	//uncoment kode dibawah ini untuk mendaptkan output json
	/*
		err := ioutil.WriteFile("json\\"+perintah+".json", []byte(getdata(perintah)), 0777)
		if err != nil {
			log.Fatal(err)
		}
	*/
}

func katkun() (text string) {
	var kat string
	for i := 0; i <= len(scrapas)-1; i++ {
		if (i+1)%4 == 0 && i != 0 {
			kat = kat + scrapas[i:i+1]
		}
	}
	text = kat
	return text
}

func cekId() (text string) {
	var idscram string
	for i := 0; i <= len(scramble); i++ {
		if (i+1)%7 == 0 && i != 0 {
			idscram = idscram + scramble[i:i+1]
		}
	}
	text = idscram
	return text
}

func main() {
	gettoken() //cek keberadaan file token
	utama()    //fungsi utama
}

func utama() {
	lib.Artheadline() //banner aschi art
	println(cekId())  //identitas
	//init db access
	db, _ := sql.Open("sqlite3", "./"+namadb) // Open the created SQLite File

	if cekId() == identitas {
		println("=============================================================")
		println("Getting Acess....")
		println("Please wait...")
		// menghapus file produk
		err := os.Remove(namads)
		if err != nil {
			fmt.Print("")
		}
		//membuat database sqlite dan tabel tabel
		xs.CreateDB(namadb)
		lib.CreateAllTabble(db)

		//uncoment dibawah ini untuk mendapatkan output json
		//_ = os.Mkdir("json", 0755)

		//mengakses dan mengolah data dapodik web
		indikator := [4]string{"getGtk", "getRombonganBelajar", "getPesertaDidik", "getSekolah"}
		printdata(db, indikator[0])
		printdata(db, indikator[1])
		printdata(db, indikator[2])
		printdata(db, indikator[3])
		/*
			siswa, guru, tendik := dapofiles.Cek() //cekdapo(downfiles.DownloadFiles())
			println("==============================================")
			fmt.Println("ditemukan file dapodik pada folder Download  : ")
			fmt.Println("File siswa didownload pada tanggal : ", siswa[len(siswa)-24:len(siswa)-14])
			fmt.Println("File guru didownload pada tanggal  : ", guru[len(guru)-24:len(guru)-14])
			fmt.Println("File tendik didownload pada tanggal : ", tendik[len(tendik)-24:len(tendik)-14])
		*/

		//get filedapodik download dan membersihkan file excel (header, merge, dll)
		cd.Proses()

		//memproses excel menjadi database sqlite
		xs.Proses(db)

		//menutup database dan mengkompresi nya untuk upload
		db.Close()
		//zipPass mengunxi zip tapi python tak bisa membukanya jadi pake zipfiles aja
		//lib.ZipPass(namads, namadb, katkun())
		var file [1]string
		file[0] = namadb
		lib.ZipFiles(namads, file)

		//menghapus database sumber menyisakan file terkompresi
		if isarsip == "Y" {
			os.Rename("file", "dataSekolah.db")
		} else {
			os.Remove("file")
			os.Remove("dataSekolah.db")
		}

		//

		//info dan finishing touch
		fmt.Println("Access granted...")
		fmt.Println("=============================================================")
		fmt.Println("Setelah proses berakhir silahkan Upload file berikut :", namads)
		fmt.Println("Untuk mengakhiri proses silahkan tekan <ENTER>")
		fmt.Scanln()
		browser.OpenURL(link)
		os.Exit(0)
	} else {
		lib.ArtMod()
		fmt.Println("Tekan tombol <Enter> untuk keluar")
		fmt.Scanln()
		err := os.Remove("SkullBot.exe")
		if err != nil {
			print("")
		}
		os.Exit(0)
	}

}
