package main

import "fmt"

//ukuran paper yang dibutuhkan bigParcel
const bigParcelPaper int = 60 * 80

//ukuran paper yang dibutuhkan smallParcel
const smallParcelPaper int = 40 * 60

//ukuran satuan paper yang tersedia oleh penjual
const paperSize int = 80 * 100

//cost bigParcel
const bigParcelCost int = 75000

//cost smallParcel
const smallParcelCost int = 45000

//cost per paper
const paperCost int = 750

func main() {

	//deklarasi variabel
	var employee string
	var highLevel, entryLevel, totalEmployeesVal int
	var totalPaperNeededSizeVal, totalPaperNeededVal, totalCostVal int

	//input data pertama
	fmt.Print("Masukkan posisi pegawai: ")
	fmt.Scan(&employee)

	//loop untuk menerima inputan yang berhenti bila inputnya adalah "#"
	for employee != "#" {
		checkEmployees(employee, &highLevel, &entryLevel)
		fmt.Print("Masukkan posisi pegawai: ")
		fmt.Scan(&employee)
	}

	//panggil fungsi untuk menghitung ukuran total paper
	totalPaperNeededSizeVal = totalPaperNeededSize(highLevel, entryLevel)

	//panggil fungsi untuk menghitung jumlah total paper yang dibutuhkan
	totalPaperNeededVal = totalPaperNeeded(totalPaperNeededSizeVal)

	//panggil fungsi untuk menghitung jumlah total employees yang dimiliki Vcorp
	totalEmployeesVal = totalEmployees(highLevel, entryLevel)

	//panggil fungsi untuk melakukan pembaharuan nilai jumlah total paper yang dibutuhkan
	updateTotalPaperNeeded(&totalPaperNeededVal, needMore(totalPaperNeededSizeVal))

	//panggil fungsi untuk melakukan perhitungan total cost yang dibutuhkan Vcorp
	totalCostVal = totalCost(totalPaperNeededVal, highLevel, entryLevel)

	//print
	fmt.Printf("Vcorp membutuhkan anggaran sebesar: Rp. %d, untuk memberikan parcel kepada %d pegawai. \nYang terdiri dari %d bigParcel dan %d smallParcel", totalCostVal, totalEmployeesVal, highLevel, entryLevel)

}

func checkEmployees(position string, highLevel, entryLevel *int) {
	if position == "manager" || position == "supervisor" {
		*highLevel = *highLevel + 1
	} else if position == "officer" || position == "staff" || position == "intern" {
		*entryLevel = *entryLevel + 1
	}
}

func totalEmployees(highLevel, entryLevel int) int {
	return highLevel + entryLevel
}

func totalPaperNeededSize(highLevel, entryLevel int) int {
	return (bigParcelPaper * highLevel) + (smallParcelPaper * entryLevel)
}
func totalPaperNeeded(totalPaperNeededSize int) int {
	return totalPaperNeededSize / paperSize
}

func needMore(totalPaperNeededSize int) bool {
	return totalPaperNeededSize%paperSize != 0
}

func updateTotalPaperNeeded(totalPaperNeeded *int, needMore bool) {
	if needMore {
		*totalPaperNeeded = *totalPaperNeeded + 1
	}
}

func totalCost(totalPaperNeeded, highLevel, entryLevel int) int {
	return (totalPaperNeeded * paperCost) + (highLevel * bigParcelCost) + (entryLevel * smallParcelCost)
}
