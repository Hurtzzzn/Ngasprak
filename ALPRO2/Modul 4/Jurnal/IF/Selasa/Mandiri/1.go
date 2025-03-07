package main

import (
	"fmt"
	"math"
)

var eloRatingKamu, eloRatingLawan float64
var skorKamu, skorLawan float64
var pertandinganKe int = 0

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			inputEloRatingKamu()
		case 2:
			inputHasilPertandingan()
		case 3:
			prosesData()
		case 4:
			cetakUpdateEloRatingKamu()
		}
		if pilih == 5 {
			break
		}
	}
}

func menu() {
	fmt.Println("-----------------------------")
	fmt.Println("           M E N U           ")
	fmt.Println("-----------------------------")
	fmt.Println("1. Input Elo Rating Awal     ")
	fmt.Println("2. Input Hasil Pertandingan  ")
	fmt.Println("3. Proses Data               ")
	fmt.Println("4. Cetak Update Elo Rating   ")
	fmt.Println("5. Exit                      ")
	fmt.Println("-----------------------------")
}

func inputEloRatingKamu() {
	fmt.Print("Masukkan elo Rating kamu: ")
	fmt.Scan(&eloRatingKamu)
}

func inputHasilPertandingan() {
	pertandinganKe++
	fmt.Printf("Pertandingan ke-%d\n", pertandinganKe)
	fmt.Print("Masukkan elo Rating lawan kamu: ")
	fmt.Scan(&eloRatingLawan)
	fmt.Print("Masukkan skor kamu (1/0.5/0): ")
	fmt.Scan(&skorKamu)
	fmt.Print("Masukkan skor lawan kamu (1/0.5/0): ")
	fmt.Scan(&skorLawan)
}

func prosesData() {
	var estimasiA float64
	const K int = 20

	estimasiA = 1.0 / (1.0 + math.Pow(10.0, (eloRatingLawan-eloRatingKamu)/400.0))

	eloRatingKamu = eloRatingKamu + float64(K)*(skorKamu-estimasiA)
	fmt.Println("Data sudah diproses")
	reset()
}

func cetakUpdateEloRatingKamu() {
	fmt.Printf("Elo rating terupdate menjadi %.2f.\n", eloRatingKamu)
}

func reset() {
	eloRatingLawan = 0
	skorKamu = 0
	skorLawan = 0
}
