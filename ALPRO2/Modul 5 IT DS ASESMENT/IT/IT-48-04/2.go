package main

import "fmt"

func barisanGeometri(n int) int {
	if n == 1 {
		return 1
	} else {
		return 3 * barisanGeometri(n-1)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(barisanGeometri(N))
}
