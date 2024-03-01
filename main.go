package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman
type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Fungsi untuk mendapatkan data teman berdasarkan nomor absen
func getDataTeman(nomorAbsen int) Teman {
	// Data teman-teman kalian (ganti dengan data teman-teman kalian sendiri)
	daftarTeman := map[int]Teman{
		1: {"Muhammad Yazid Fauzan", "Lampung", "Freelancer", "Menambah pengetahuan tentang bahasa pemrograman baru, untuk membuka kesempatan baru di dunia kerja programming"},
		2: {"Riezki Riedo Tama", "Bandung", "Pegawai", "Ingin mempelajari hal baru yaitu tentang programming"},
		3: {"Muhammad Rizky Pratama", "Tanggerang", "Mahasiswa", "Ingin tahu lebih tentang bahasa pemrograman yang sedang trend"},
		4: {"Irfan Iqbal", "Cikarang", "Mahasiswa", "Rasa ingin tahu tentang dunia programming yang mendorong saya untuk ikut"},
		5: {"Jojo Refrendo", "Lampung", "Pegawai", "Menambah pengetahuan baru tentang programming"},
		
	}

	// Mengambil data teman sesuai nomor absen
	teman, ok := daftarTeman[nomorAbsen]
	if !ok {
		// Jika nomor absen tidak ditemukan, return data kosong
		return Teman{}
	}
	return teman
}

func main() {
	// Memeriksa apakah argumen yang diberikan valid
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go [nomor_absen]")
		return
	}

	// Mengambil nomor absen dari argumen
	nomorAbsen := os.Args[1]

	// Mengonversi nomor absen dari string ke integer
	var nomorAbsenInt int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &nomorAbsenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka.")
		return
	}

	// Mendapatkan data teman berdasarkan nomor absen
	teman := getDataTeman(nomorAbsenInt)

	// Menampilkan data teman
	if teman.Nama != "" {
		fmt.Println("Nama:", teman.Nama)
		fmt.Println("Alamat:", teman.Alamat)
		fmt.Println("Pekerjaan:", teman.Pekerjaan)
		fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
	} else {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan.")
	}
}
