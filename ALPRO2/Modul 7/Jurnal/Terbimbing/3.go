package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var n int
	baca(&data, &n)
	cetak(data, n)
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array A dan banyak elemen n terdefinisi sembarang
		FS: Array A berisi sebanyak n elemen dengan aturan pengisian array sebagai berikut:
			Array A hanya akan diisi, hanya jika indeks saat ini genap, yaitu: 0, 2, 4, ... dan
			bilangan yang diinput genap, atau hanya jika indeks saat ini ganjil, yaitu: 1, 3, 5, ...
			dan bilangan yang diinput ganjil.
			Program tidak memperkenankan pengisian bilangan 0 atau -1. Jika diisi dengan bilangan
			0 atau -1, program harus meminta pengisian baru.
	*/
	var x, i int
	i = 0
	fmt.Scan(&x)
	if x == 0 || x == -1 {
		fmt.Scan(&x)
	}
	for i < NMAX {
		if i%2 == 0 && x%2 == 0 {
			A[i] = x
			i++
		} else if i%2 != 0 && x%2 != 0 {
			A[i] = x
			i++
		}
		fmt.Scan(&x)
		if x == 0 || x == -1 {
			fmt.Scan(&x)
		}
	}
	*n = i
}

func cetak(A tabInt, n int) {
	/*
		IS: Array A dengan n elemen terdefinisi
		FS: Array A dengan n elemen tercetak di layar
	*/
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}
