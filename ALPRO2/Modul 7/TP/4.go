package main

import "fmt"

const MAXNUM int = 10

type tabInt [MAXNUM]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array A dan n terdefinisi sembarang
		Proses: Setiap bilangan yang dibaca ditampung dalam sebuah variabel.
		    Jika bilangan negatif, maka ubah menjadi bilangan positif dan
		    masukan ke dalam array. Pembacaan berakhir jika terbaca bilangan 0.
		FS: Array A sebanyak n elemen berisi bilangan positif
	*/
	var x, i int
	fmt.Scan(&x)
	i = 0
	for x != 0 && i < MAXNUM {
		if x < 0 {
			x = x * -1
		}
		A[i] = x
		i++
		fmt.Scan(&x)
	}
	*n = i
}

func cetak(A tabInt, n int) {
	/*
		IS: Array A terdefinisi sebanya kn elemen
		FS: Array A tercetak di layar
	*/
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	/* Mengembalikan jumlah dari nilai bilangan pada elemen array A */
	var i, jumlah int
	jumlah = 0
	for i = 0; i < n; i++ {
		jumlah += A[i]
	}
	return jumlah
}

func rata_rata(A tabInt, n int) float64 {
	/* Mengembalikan rata-rata bilangan */
	var rata float64
	rata = float64(jumlah(A, n)) / float64(n)
	return rata
}
