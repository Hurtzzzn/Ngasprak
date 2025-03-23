package main

import "fmt"

func barisanSegitiga(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + barisanSegitiga(n-1)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(barisanSegitiga(N))
}
