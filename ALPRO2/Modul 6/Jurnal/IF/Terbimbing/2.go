package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	cetakBaris(n)

}

func cetakBaris(n int) {
	if n == 1 {
		fmt.Print(1)
	} else {
		if n%2 == 0 {
			cetakBaris(n - 1)
			fmt.Print(" + ", n)
		} else {
			cetakBaris(n - 1)
			fmt.Print(" - ", n)
		}

	}
}
