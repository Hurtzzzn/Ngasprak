package main

import "fmt"

const NMAX int = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func main() {
	var nilaiAkhir tabInt
	bacaData(&nilaiAkhir)
	bacaData(&nilaiAkhir)
	bacaData(&nilaiAkhir)
	bacaData(&nilaiAkhir)
	bacaData(&nilaiAkhir)
	bacaData(&nilaiAkhir)
	cetakData(nilaiAkhir)
}

func bacaData(A *tabInt) {
	/*
		  IS: A.info[i] adalah field yang menampung data array, sedangkan A.n adalah field untuk menampung banyaknya
			  elemen data. Kedua field itu terdefinisi sembarang yang berarti bisa kosong atau berisi nilai.
		  FS: Data array A.info bertambah satu elemen dan A.n bertambah satu nilai. keduanya bertambah, selama belum
			  melebihi kapasitas maksimum array (NMAX)
	*/
	var x int
	fmt.Scan(&x)
	if x > 0 && A.n < NMAX {
		A.info[A.n] = x
		A.n++
	}
}

func cetakData(A tabInt) {
	/*
			IS: Array A terdefinisi sembarang, yang berarti bisa berisi nilai atau kosong
		  	FS: Mencetak sebanyak A.n elemen array dengan format:
		       <e0> <e1> <e2> ... <en-1>".
	*/
	var i int
	for i = 0; i < A.n; i++ {
		fmt.Printf("%d ", A.info[i])
	}
	fmt.Println()
}
