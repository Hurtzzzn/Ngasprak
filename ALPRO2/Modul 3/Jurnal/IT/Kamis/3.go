package main

import "fmt"

var jenisKendaraan string
var lamaRencanaPinjam, lamaKenyataanPinjam int
var totalPemasukan int

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			inputPinjamKendaraan()
		case 2:
			inputKembaliKendaraan()
		case 3:
			hitungBiayaRental()
		case 4:
			cetakTotalPemasukan()
		}
		if pilih == 5 {
			break
		}
	}
}

func menu() {
	fmt.Println("-----------------------------")
	fmt.Println("           M E N U        ")
	fmt.Println("-----------------------------")
	fmt.Println("1. Input Pinjam Kendaraan")
	fmt.Println("2. Input Kembali Kendaraan")
	fmt.Println("3. Hitung Biaya Rental")
	fmt.Println("4. Cetak Total Pemasukan")
	fmt.Println("5. Exit")
	fmt.Println("-----------------------------")
}

func inputPinjamKendaraan() {
	fmt.Print("Masukkan jenis kendaraan yang dipinjam (minibus/sedan): ")
	fmt.Scan(&jenisKendaraan)
	fmt.Print("Masukkan lama hari peminjaman: ")
	fmt.Scan(&lamaRencanaPinjam)
}

func inputKembaliKendaraan() {
	fmt.Print("Masukkan lama hari peminjaman: ")
	fmt.Scan(&lamaKenyataanPinjam)
}

func hitungBiayaRental() {
	var denda, hargaSewa, totalBiayaRental int
	if jenisKendaraan == "minibus" {
		hargaSewa = 400000
	} else if jenisKendaraan == "sedan" {
		hargaSewa = 300000
	}
	if lamaKenyataanPinjam > lamaRencanaPinjam {
		denda = (lamaKenyataanPinjam - lamaRencanaPinjam) * hargaSewa * 2
	}

	totalBiayaRental = lamaRencanaPinjam*hargaSewa + denda
	fmt.Printf("Biaya rental %s selama %d hari: Rp %d,-\n", jenisKendaraan, lamaKenyataanPinjam, totalBiayaRental)
	totalPemasukan += totalBiayaRental
	reset()
}

func reset() {
	jenisKendaraan = ""
	lamaRencanaPinjam = 0
	lamaKenyataanPinjam = 0
}

func cetakTotalPemasukan() {
	fmt.Printf("Total uang: %d\n", totalPemasukan)
}
