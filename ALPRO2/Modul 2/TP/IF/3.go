//WAREHOUSE

//IMPORTANT:
//Disarankan Bentuk Pengerjaan Soal Ini Ada Clue Mirip Answer Box Preload Kalau di LMS jadi Mahasiswa Melengkapi Template Code

package main

import (
	"fmt"
)

func main() {

	// struct product
	type Product struct {
		name  string
		price float64
		stock int
	}

	var p Product
	var status string
	var priceDiskon float64

	// Input Data Produk dari Pengguna
	fmt.Print("Masukkan nama Produk: ")
	fmt.Scanln(&p.name)

	for p.name != "#" {

		fmt.Print("Masukkan harga: ")
		fmt.Scanln(&p.price)

		fmt.Print("Masukkan stok: ")
		fmt.Scanln(&p.stock)

		// Menentukan Status berdasarkan Stok
		status = "Produk habis"
		if p.stock > 0 {
			status = "Produk tersedia"
		}

		// Menentukan Diskon
		var diskonPersen float64
		if p.price >= 1000000 {
			diskonPersen = 10
		} else if p.price >= 500000 {
			diskonPersen = 5
		} else {
			diskonPersen = 0
		}

		priceDiskon = p.price - (p.price * diskonPersen / 100)

		// Menampilkan Output
		fmt.Println("\nDetail Produk:")
		fmt.Println("Nama:", p.name)
		fmt.Printf("Harga Asli: Rp.%.0f\n", p.price)
		fmt.Printf("Diskon: %.0f%%\n", diskonPersen)
		fmt.Printf("Harga Setelah Diskon: Rp.%.0f\n", priceDiskon)
		fmt.Println("Stok:", p.stock)
		fmt.Println("Status:", status)

		// Konfirmasi Loop
		fmt.Print("\nMasukkan name Produk: ")
		fmt.Scanln(&p.name)
	}
}
