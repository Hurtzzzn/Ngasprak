package main

import "fmt"

var gol, kegolan int
var jumMenang, jumDraw, jumKalah, jumGol, jumKegol int
var jumSelisihGol, jumPoint int

func main() {
	var n, i int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&gol, &kegolan)
		hitungMenang(gol, kegolan, &jumMenang)
		hitungDraw(gol, kegolan, &jumDraw)
		hitungKalah(gol, kegolan, &jumKalah)
		hitungJumGolKegolanSelisih(gol, kegolan, &jumGol, &jumKegol, &jumSelisihGol)
	}
	hitungJumPoint(&jumPoint)
	fmt.Println(n, jumMenang, jumDraw, jumKalah, jumGol, jumKegol, jumSelisihGol, jumPoint)
}

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg = *jg + g
	*jk = *jk + k
	*jsg = *jg - *jk
}

func hitungJumPoint(jp *int) {
	*jp = jumMenang*3 + jumDraw*1
}
