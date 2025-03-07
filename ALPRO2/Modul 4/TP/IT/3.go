package main

import "fmt"

// iteratif
func faktorialIter(n int) int {
	var fact, i int
	fact = 1
	for i = n; i > 0; i-- {
		fact = fact * i
	}
	return fact
}

// rekursif
func faktorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(faktorial(num))
	//fmt.Println(faktorialIter(num))
}
