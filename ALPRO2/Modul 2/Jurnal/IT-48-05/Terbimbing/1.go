package main

import "fmt"

func vowel_uppercase(kar byte) bool {
	var vowel_upper_case bool
	vowel_upper_case = kar == 'A' || kar == 'I' || kar == 'U' || kar == 'E' || kar == 'O'
	return vowel_upper_case
}

func main() {
	var abjad byte
	var jumlah int
	fmt.Scanf("%c", &abjad)
	jumlah = 0
	for abjad != '.' {
		if vowel_uppercase(abjad) {
			jumlah++
		}
		fmt.Scanf("%c", &abjad)
	}
	fmt.Printf("Jumlah vokal uppercase adalah %d\n", jumlah)
}
