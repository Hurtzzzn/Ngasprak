package main

import "fmt"

func ganjil(a, b, c, d, e int) bool {
	// Mengembalikan true apabila masing-masing a, b, c, d, dan e adalah ganjil
	return a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0
}

func genap(a, b, c, d, e int) bool {
	// Mengembalikan true apabila masing-masing a, b, c, d, dan e adalah genap
	return a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0
}

func diskon(member bool, c, d, e int) float64 {
	// Mengembalikan nilai diskon sebesar 1.5x dari total c, d, dan e.
	// Tambahkan 15% dari nilai diskon awal apabila member
	// Nilai diskon maksimal adalah 50%
	var disc float64
	disc = 1.5 * float64(c+d+e)
	if member {
		disc = disc + 0.15*disc
	}
	if disc > 50 {
		disc = 50
	}
	return disc
}

func cashback(member bool, a, b, c int) float64 {
	// Mengembalikan nilai cashback sebesar 3.5x dari total a, b, dan c.
	// Tambahkan 15% dari nilai cashback awal apabila member
	// Nilai cashback maksimal adalah 35%
	var cash float64

	cash = 3.5 * float64(a+b+c)
	if member {
		cash = cash + 0.15*cash
	}
	if cash > 35 {
		cash = 35
	}
	return cash
}

func main() {
	var membership bool
	var p1, p2, p3, p4, p5 int
	var persenDiskon, persenCashback float64

	fmt.Scan(&membership, &p1, &p2, &p3, &p4, &p5)

	if !ganjil(p1, p2, p3, p4, p5) && !genap(p1, p2, p3, p4, p5) {
		persenDiskon = diskon(membership, p3, p4, p5)
		persenCashback = cashback(membership, p1, p2, p3)
	} else if ganjil(p1, p2, p3, p4, p5) {
		persenDiskon = diskon(membership, p3, p4, p5)
	} else if genap(p1, p2, p3, p4, p5) {
		persenCashback = cashback(membership, p1, p2, p3)
	}

	fmt.Printf("%.2f %.2f\n", persenDiskon, persenCashback)
}
