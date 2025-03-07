package main

import "fmt"

func power(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * power(a, b-1)
}

func main() {
	var b, e int
	fmt.Scan(&b, &e)
	fmt.Println(power(b, e))
}
