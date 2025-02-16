package main

import "fmt"

func main() {
	var X, i int
	var prima bool
	fmt.Scan(&X)
	prima = X > 1
	for i = 2; i < X; i++ {
		prima = prima && (X%i != 0)
	}
	if prima {
		fmt.Println("Prima")
	} else {
		fmt.Println("Bukan Prima")
	}
}
