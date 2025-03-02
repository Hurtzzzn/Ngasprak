package main

import "fmt"

func menu() {
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan ")
	fmt.Println("2. Hitung Perkalian   ")
	fmt.Println("3. Hitung Pembagian   ")
	fmt.Println("4. Exit   ")
	fmt.Println("-----------------------")
}

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			hitungJumlah()
		case 2:
			hitungKali()
		case 3:
			hitungBagi()
		}
		if pilih == 4 {
			break
		}
	}
}

func hitungJumlah() {
	var b1, b2, hasil int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&b1, &b2)
	hasil = b1 + b2
	fmt.Println("Hasil penjumlahan:", hasil)
}

func hitungKali() {
	var b1, b2, hasil int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&b1, &b2)
	hasil = b1 * b2
	fmt.Println("Hasil perkalian:", hasil)
}

func hitungBagi() {
	var b1, b2, hasil float64
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&b1, &b2)
	hasil = b1 / b2
	fmt.Println("Hasil pembagian:", hasil)
}
