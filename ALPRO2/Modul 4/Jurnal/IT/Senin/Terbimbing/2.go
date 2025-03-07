package main

import "fmt"

// Iteratif
func jumlahDigitIter(n int) int {
	var jum int
	jum = 0
	for n > 0 {
		jum = jum + n%10
		n = n / 10
	}
	return jum
}

// rekursif
func jumlahDigit(n int) int {
	if n == 0 {
		return 0
	} else {
		return (n % 10) + jumlahDigit(n/10)
	}
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(jumlahDigit(num))
	// fmt.Println(jumlahDigitIter(num))
}
