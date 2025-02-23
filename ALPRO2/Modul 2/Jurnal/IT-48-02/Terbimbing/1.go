package main

import "fmt"

func consonant(kar byte) bool {
	var not_vowel_lower_case, not_vowel_upper_case bool
	not_vowel_lower_case = kar != 'a' && kar != 'i' && kar != 'u' && kar != 'e' && kar != 'o'
	not_vowel_upper_case = kar != 'A' && kar != 'I' && kar != 'U' && kar != 'E' && kar != 'O'
	return (not_vowel_lower_case && not_vowel_upper_case)
}

func main() {
	var abjad byte
	var jumlah int
	fmt.Scanf("%c", &abjad)
	jumlah = 0
	for abjad != '.' {
		if consonant(abjad) {
			jumlah++
		}
		fmt.Scanf("%c", &abjad)
	}
	fmt.Printf("Jumlah konsonan adalah %d\n", jumlah)
}
