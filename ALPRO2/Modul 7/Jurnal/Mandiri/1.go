package main

import "fmt"

const NMAX int = 20

type tabInt [NMAX]int

func main() {
	var dataAwal, dataHasil tabInt
	var nAwal, nHasil int
	var a, b float64
	baca(&dataAwal, &nAwal)
	fmt.Scan(&a, &b)
	jumlah_nilai(dataAwal, &dataHasil, nAwal, &nHasil, a, b)
	cetak(dataHasil, nHasil)
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array (A) dan banyak elemen (n) terdefinisi sembarang
		FS: Array (A) berisi bilangan bulat sejumlah n elemen
	*/
	var i int
	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func jumlah_nilai(A tabInt, B *tabInt, nA int, nB *int, a, b float64) {
	/*
		IS: Array A sebanyak nA elemen terdefinisi. Batas bawah a dan batas atas b terdefinisi
		FS: Array B sebanyak nB elemen berisi nilai
	*/
	var i, k int
	k = 0
	for i = 0; i < nA; i++ {
		if float64(A[i]) >= a && float64(A[i]) <= b {
			*nB++
			B[k] = A[i]
			k++
		}
	}
}

func cetak(B tabInt, n int) {
	/*
		IS: Array B sebanyak n terdefinisi
		FS: Mencetak elemen B sebanyak n elemen dengan format:
			Ada <n> nilai: <b1 b2 b3 ... n>
	*/
	var i int
	fmt.Printf("Ada %d nilai: ", n)
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", B[i])
	}
	fmt.Println()
}
