package main

// Formasi Lucas 2 1 3 4 7 11 18 29 ...

import "fmt"

func main() {
	var suku, bil_lucas int
	fmt.Scan(&suku)
	bil_lucas = lucas(suku)
	fmt.Println(bil_lucas)
}

func lucas(n int) int {
	if n == 1 {
		return 2
	} else if n == 2 {
		return 1
	} else if n > 2 {
		return lucas(n-1) + lucas(n-2)
	}
	return 0
}
