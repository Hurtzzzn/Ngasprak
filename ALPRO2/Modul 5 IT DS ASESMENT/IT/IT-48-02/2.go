package main

import "fmt"

func barisanGanjil(n int) int {
	if n == 1 {
		return 1
	} else {
		return 2 + barisanGanjil(n-1)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(barisanGanjil(N))
}
