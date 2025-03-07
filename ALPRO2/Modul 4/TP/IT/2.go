package main

import "fmt"

func persamaan2(a, b int) int {
	if b == 0 {
		return a
	} else {
		return persamaan2(b, a%b)
	}
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(persamaan2(x, y))
}
