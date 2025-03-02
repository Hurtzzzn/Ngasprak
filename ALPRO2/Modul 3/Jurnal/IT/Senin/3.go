package main

import "fmt"

// variabel global
var jenisKendaraan string
var jam1, menit1, detik1 int
var jam2, menit2, detik2 int
var totalUang int

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			inputKendaraanMasuk()
		case 2:
			inputKendaraanKeluar()
		case 3:
			hitungBiayaParkir()
		case 4:
			cetakTotalUang()
		}
		if pilih == 5 {
			break
		}
	}
}

func menu() {
	fmt.Println("--------------------------")
	fmt.Println("          M E N U        ")
	fmt.Println("--------------------------")
	fmt.Println("1. Input Kendaraan Masuk")
	fmt.Println("2. Input Kendaraan Keluar")
	fmt.Println("3. Hitung Biaya Parkir")
	fmt.Println("4. Cetak Total Uang")
	fmt.Println("5. Exit")
	fmt.Println("-------------------------")
}

func inputKendaraanMasuk() {
	fmt.Print("Masukkan jenis kendaraan (mobil/motor): ")
	fmt.Scan(&jenisKendaraan)
	fmt.Print("Masukkan jam, menit, detik kendaraan masuk: ")
	fmt.Scan(&jam1, &menit1, &detik1)
}

func inputKendaraanKeluar() {
	fmt.Print("Masukkan jam, menit, detik kendaraan keluar: ")
	fmt.Scan(&jam2, &menit2, &detik2)
}

func hitungBiayaParkir() {
	var totalDetik, durasiJam, biayaParkir int
	totalDetik = (jam2*3600 + menit2*60 + detik2) - (jam1*3600 + menit1*60 + detik1)
	durasiJam = totalDetik / 3600
	if totalDetik%3600 > 0 {
		durasiJam++
	}

	if jenisKendaraan == "mobil" {
		biayaParkir = 5000 + (durasiJam-1)*3000
	} else if jenisKendaraan == "motor" {
		biayaParkir = 2000 + (durasiJam-1)*1000
	}
	fmt.Printf("Biaya parkir %s selama %d jam: Rp %d,-\n", jenisKendaraan, durasiJam, biayaParkir)
	totalUang += biayaParkir
	reset()
}

func reset() {
	jenisKendaraan = ""
	jam1 = 0
	menit1 = 0
	detik1 = 0
	jam2 = 0
	menit2 = 0
	detik2 = 0
}

func cetakTotalUang() {
	fmt.Printf("Total uang: Rp %d,-\n", totalUang)
}
