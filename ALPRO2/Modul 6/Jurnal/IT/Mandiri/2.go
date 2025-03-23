package main

import (
	"fmt"
	"math"
)

 Tipe bentukan struct titik dengan field x dan y bertipe float64
type titik struct {
	x, y float64
}

 Tipe bentukan struct garis dengan field t1 dan t2 bertipe titik
 serta panjang bertipe float64
type garis struct {
	t1, t2  titik
	panjang float64
}

func main() {
	var p1, p2, p3, p4 titik  //p = point
	var l1, l2 garis          //l = line
	var x, y float64

	baca_x_y(&x, &y)
	p1 = titik_baru(x, y)

	baca_x_y(&x, &y)
	p2 = titik_baru(x, y)

	baca_x_y(&x, &y)
	p3 = titik_baru(x, y)

	baca_x_y(&x, &y)
	p4 = titik_baru(x, y)

	l1 = garis_baru(p1, p2)
	hitung_panjang_garis(&l1)

	l2 = garis_baru(p3, p4)
	hitung_panjang_garis(&l2)

	cetak_data(l1)
	cetak_data(l2)
}

func baca_x_y(a, b *float64) {
	/* IS: a dan b terdefinisi sembarang (data a dan b tersedia pada piranti masukan)
	   FS: a dan b berisi nilai
	*/
	fmt.Scan(a, b)
}

func titik_baru(a, b float64) titik {
	/* Mengembalikan titik dengan field a dan b */
	var p titik
	p.x = a
	p.y = b
	return p
}

func garis_baru(a, b titik) garis {
	/* Mengembalikan garis dengan field titik a dan b */
	var g garis
	g.t1 = a
	g.t2 = b
	return g
}

func hitung_panjang_garis(g *garis) {
	/* Mengembalikan panjang g.panjang yang dihitung dengan memanggil fungsi panjang */
	g.panjang = panjang(g.t1, g.t2)
}

func panjang(a, b titik) float64 {
	/* Mengembalikan panjang atau jarak dari titik a ke b dengan menggunakan fungsi
	   math.Sqrt (akar) dan math.Pow (pangkat) */
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func cetak_data(g garis) {
	/* IS: Garis g terdefinisi
	   FS: Tercetak data garis g dengan format:
	       "Garis dibentuk oleh titik <x, y)> dan <x, y> memiliki panjang sebesar <panjang>."
	*/
	fmt.Printf("Garis dibentuk oleh titik (%v, %v) dan (%v, %v) ", g.t1.x, g.t1.y, g.t2.x, g.t2.y)
	fmt.Printf("memiliki panjang sebesar %v.\n", g.panjang)
}
