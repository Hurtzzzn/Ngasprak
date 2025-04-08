package main

import "fmt"

func main() {
	const MAXDATA int = 100
	var A [MAXDATA]int
	var i, n, banyakElemen, jumlah int
	var rataRata float64

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	jumlah = 0
	for i = 0; i < n; i++ {
		jumlah += A[i]
	}
	rataRata = float64(jumlah) / float64(n)

	banyakElemen = 0
	for i = 0; i < n; i++ {
		if float64(A[i]) < rataRata {
			banyakElemen++
		}
	}
	fmt.Println(rataRata, banyakElemen)
}
