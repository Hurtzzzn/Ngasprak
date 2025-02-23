//VERIFIED CREATOR

package main

import "fmt"

func main() {

	//Struct Creator
	type Creator struct {
		name            string
		monthlyListener int
		numberOfSongs   int
	}

	//Deklarasi Variabel
	var c Creator
	var diffM, diffN int

	//Inputs
	fmt.Scan(&c.name)
	fmt.Scan(&c.monthlyListener)
	fmt.Scan(&c.numberOfSongs)

	//Hitung Selisih
	diffM = 500000 - c.monthlyListener
	diffN = 10 - c.numberOfSongs

	//Conditional Printing
	if c.numberOfSongs >= 10 && c.monthlyListener >= 500000 {
		fmt.Printf("Kreator %s dapat terverifikasi dengan\njumlah total lagu: %d\npendengar bulanan sebanyak: %d\n", c.name, c.numberOfSongs, c.monthlyListener)
	} else if c.numberOfSongs < 10 && c.monthlyListener >= 500000 {
		fmt.Printf("Kreator %s belum dapat terverifikasi, butuh %d lagu rilis lagi\n", c.name, diffN)
	} else if c.numberOfSongs >= 10 && c.monthlyListener < 500000 {
		fmt.Printf("Kreator %s belum dapat terverifikasi, butuh %d pendengar bulanan lagi\n", c.name, diffM)
	} else {
		fmt.Printf("Kreator %s belum dapat terverifikasi butuh\njumlah total lagu: %d lagi\npendengar bulanan sebanyak: %d lagi\n", c.name, diffN, diffM)
	}

}
