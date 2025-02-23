package main

import "fmt"

func fungsi_f(x, y, z float64) float64 {
	return 2 * x / (x + y + z)
}

func fungsi_g(x, y float64) float64 {
	return 1/5*x + y
}

func main() {
	var bil_1, bil_2, bil_3, hasil_1, hasil_2 float64

	fmt.Scan(&bil_1, &bil_2, &bil_3)
	hasil_1 = fungsi_f(bil_1, bil_2, bil_3)
	hasil_2 = fungsi_g(bil_1, bil_2)
	fmt.Println(hasil_1, hasil_2)
}
