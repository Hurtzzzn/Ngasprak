package main

import "fmt"

const NMAX int = 10

type tim struct {
	gol                                    [NMAX]int
	totPertandingan, totGol, totKebobolan  int
	totMenang, totKalah, totDraw, totPoint int
}

func main() {
	var timA, timB tim
	bacaHasil(&timA, &timB)
	fmt.Println("Statistik Tim A")
	cetakHasil(timA)
	fmt.Println("\nStatistik Tim B")
	cetakHasil(timB)
}

func bacaHasil(tA, tB *tim) {
	/*
		IS: Tim A (tA) dan tim B (tB) terdefinisi sembarang
		Proses: Membaca skor pertandingan berupa jumlah gol tim A dan tim B.
				Mengisi field-field sesuai skor kedua tim. Field-fieldnya:
				1. total pertandingan
				2. gol setiap pertandingan
				3. total gol
				4. total kebobolan
				5. total point: point 3 untuk menang, point 1 untuk draw
				6. total menang: Menang, jika gol lebih besar dari gol lawan
				7. total draw: Draw, jika gol sama dengan gol lawan
				8. total kalah: Kalah, jika gol lebih kecil dari gol lawan
				Pembacaan skor berakhir jika kedua skor bernilai negatif.
		FS: Data tim A (tA) dan data tim B (tB) berisi nilai
	*/
	var i int
	var x, y int
	fmt.Scan(&x, &y)
	i = 0
	for x >= 0 && y >= 0 && i < NMAX {
		tA.gol[i] = x
		tB.gol[i] = y
		tA.totGol += x
		tA.totKebobolan += y
		tB.totGol += y
		tB.totKebobolan += x
		if x > y {
			tA.totPoint += 3
			tA.totMenang++
			tB.totKalah++
		} else if x < y {
			tB.totPoint += 3
			tB.totMenang++
			tA.totKalah++
		} else {
			tA.totPoint++
			tB.totPoint++
			tA.totDraw++
			tB.totDraw++
		}
		i++
		fmt.Scan(&x, &y)
	}
	tA.totPertandingan = i
	tB.totPertandingan = i
}

func cetakHasil(t tim) {
	/*
		IS: t terdefinisi
		FS: tercetak di layar statistik pertandingan dengan format:

		Total pertandingan: <total pertandingan>
		Gol tiap pertandingan: <gol1 gol2 ... goln>
		Total menang: <total kemenangan>
		Total draw: <total draw>
		Total kalah: <total kalah>
		Total gol: <total gol>
		Total kebobolan: <total kebobolan>
		Total point: <total point>
	*/
	var i, n int
	n = t.totPertandingan
	fmt.Println("Total pertandingan:", n)
	fmt.Print("Gol tiap pertandingan: ")
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", t.gol[i])
	}
	fmt.Println()
	fmt.Println("Total menang:", t.totMenang)
	fmt.Println("Total draw:", t.totDraw)
	fmt.Println("Total kalah:", t.totKalah)
	fmt.Println("Total gol:", t.totGol)
	fmt.Println("Total kebobolan:", t.totKebobolan)
	fmt.Println("Total point:", t.totPoint)
}
