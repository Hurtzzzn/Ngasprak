package main

import "fmt"

func adaBilanganM(n, m int) string {
	var digit int
	for n > 0 {
		digit = n % 10
		if digit == m {
			return "YA"
		}
		n /= 10
	}
	return "TIDAK"
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	if N < 10 || N > 999999 {
		return
	}

	fmt.Println(adaBilanganM(N, M))
}
