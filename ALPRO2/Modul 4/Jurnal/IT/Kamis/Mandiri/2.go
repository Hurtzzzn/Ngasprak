package main

import "fmt"

func persamaan(n int) int {
	if n == 2 {
		return 0
	}
	return persamaan(n-2) + n
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(persamaan(num))
}
