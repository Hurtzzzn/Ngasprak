//FINANCE

package main

import "fmt"

func main() {

	type Account struct {
		name            string
		balance         float64
		lastMonthIncome float64
		goal            float64
	}

	var a Account
	var goalVal float64
	var savings float64

	fmt.Print("Nama kamu: ")
	fmt.Scan(&a.name)

	fmt.Print("Masukkan harapan jumlah tabungan kamu setiap bulannya (%): ")
	fmt.Scan(&a.goal)

	fmt.Print("Masukkan pendapatan bulan lalu kamu: ")
	fmt.Scan(&a.lastMonthIncome)

	fmt.Print("Masukkan pendapatan kamu yang masih tersisa dari bulan lalu: ")
	fmt.Scan(&a.balance)

	goalVal = a.lastMonthIncome * a.goal / 100
	savings = a.lastMonthIncome - a.balance

	//Main Logic
	if savings == goalVal {
		fmt.Printf("%s, kamu berhasil menabung sesuai target, %.0f%% dari pendapatan kamu\n", a.name, a.goal)
	} else if savings < goalVal {
		fmt.Printf("%s, kamu berhasil menabung melebihi target, %.0f%% dari pendapatan kamu\n", a.name, a.goal)
	} else {
		fmt.Printf("%s, kamu gagal menabung sesuai target, %.0f%% dari pendapatan kamu\n", a.name, a.goal)
	}
}
