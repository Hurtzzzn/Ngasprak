package main

import "fmt"

func main() {
	var N, temp, f1, f2, ganjil int
	fmt.Scan(&N)
	f1 = 1
	f2 = 1
	for ganjil < N {
		if f1%2 != 0 {
			fmt.Print(f1, " ")
			ganjil++
		}
		temp = f2
		f2 = f1 + f2
		f1 = temp
	}
}
