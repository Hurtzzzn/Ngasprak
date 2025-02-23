package main

import "fmt"

func main() {
	var karakter, hasil byte
	fmt.Scanf("%c ", &karakter)
	hasil = lowToUpper(karakter)
	fmt.Printf("%c\n", hasil)
}

func lowToUpper(kar byte) byte {
	return kar - 32
}
