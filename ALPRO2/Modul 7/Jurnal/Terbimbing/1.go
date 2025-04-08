package main

import "fmt"

func main() {
	const MAXDATA int = 100
	var A [MAXDATA]int
	var i, n, banyakElemen, jumlah int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	banyakElemen = 0
	jumlah = 0
	for i = 0; i < n; i++ {
		if A[i]%2 == 0 {
			banyakElemen++
			jumlah += A[i]
		}
	}
	fmt.Println(banyakElemen, jumlah)
}
