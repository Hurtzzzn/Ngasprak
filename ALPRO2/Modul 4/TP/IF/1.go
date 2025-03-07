package main

import "fmt"

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	const pi float64 = 3.14
	*l = pi * r * r
	*k = 2 * pi * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL, lP, kL, kP float64, totLuas, totKel *float64) {
	*totLuas = lL + lP
	*totKel = kL + kP
}

func main() {
	var radius, sisi float64
	var luasLingkaran, kelilingLingkaran, luasPersegi, kelilingPersegi float64
	var totalLuas, totalKeliling float64

	fmt.Scan(&radius, &sisi)
	if radius != 0 && radius != 0 {
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	}
	for radius != 0 && sisi != 0 {
		hitungLuasKelilingLingkaran(radius, &luasLingkaran, &kelilingLingkaran)
		hitungLuasKelilingPersegi(sisi, &luasPersegi, &kelilingPersegi)
		hitungTotal(luasLingkaran, luasPersegi, kelilingLingkaran, kelilingPersegi, &totalLuas, &totalKeliling)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", radius, sisi, luasLingkaran, luasPersegi, kelilingLingkaran, kelilingPersegi, totalLuas, totalKeliling)
		fmt.Scan(&radius, &sisi)
	}
}
