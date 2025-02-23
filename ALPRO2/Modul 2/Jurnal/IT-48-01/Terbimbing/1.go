package main

import "fmt"

func vowel(kar byte) bool {
	var vowel_lower_case, vowel_upper_case bool
	vowel_lower_case = kar == 'a' || kar == 'i' || kar == 'u' || kar == 'e' || kar == 'o'
	vowel_upper_case = kar == 'A' || kar == 'I' || kar == 'U' || kar == 'E' || kar == 'O'
	return vowel_lower_case || vowel_upper_case
}

func main() {
	var abjad byte
	var jumlah int
	fmt.Scanf("%c", &abjad)
	jumlah = 0
	for abjad != '.' {
		if vowel(abjad) {
			jumlah++
		}
		fmt.Scanf("%c", &abjad)
	}
	fmt.Printf("Jumlah vokal adalah %d\n", jumlah)
}
