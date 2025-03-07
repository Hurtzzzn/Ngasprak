package main

import "fmt"

func balikAngka(n, balik int) int {
	if n == 0 {
		return balik
	}
	return balikAngka(n/10, balik*10+n%10)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(balikAngka(N, 0))
}
