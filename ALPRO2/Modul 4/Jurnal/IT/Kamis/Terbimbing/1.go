package main

import "fmt"

func persamaan(n int) int {
	if n == 0 {
		return 1
	}
	return persamaan(n-1) * n
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(persamaan(num))
}
