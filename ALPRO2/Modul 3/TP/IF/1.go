package main

import "fmt"

func main() {
	var CeAwal, CeAkhir, step, Re, Fa, Ke float64
	var iterator float64
	fmt.Scan(&CeAwal, &CeAkhir, &step)
	fmt.Printf("%10s %10s %10s %10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")
	for iterator = CeAwal; iterator <= CeAkhir; iterator += step {
		Re = reamur(iterator)
		Fa = fahrenheit(iterator)
		Ke = kelvin(iterator)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", iterator, Re, Fa, Ke)
	}
}

func reamur(C float64) float64 {
	return 4.0 / 5.0 * C
}

func fahrenheit(C float64) float64 {
	return 9.0/5.0*C + 32
}

func kelvin(C float64) float64 {
	return 273 + C
}
