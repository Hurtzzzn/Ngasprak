package main

import "fmt"

func main() {
	var totalValid, totalTidakValid, i, n, voucher int
	fmt.Scan(&n)
	totalValid = 0
	totalTidakValid = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&voucher)
		if valid(voucher) {
			totalValid++
		} else {
			totalTidakValid++
		}
	}
	fmt.Println(totalValid, totalTidakValid)
}

func len(num int) int {
	/* mengembalikan banyaknya digit dari bilangan bulat num */
	var i int = 0
	for num > 0 {
		num = num / 10
		i += 1
	}
	return i
}

func digitAwalAkhir(num int, d1, d2, d3, d4 *int) {
	/* I.S. terdefinisi bilangan bulat num yang menyatakan 5 atau 6 digit nomor seri voucher
	   F.S. d1 dan d2 berisi digit pertama dan kedua (dari kiri) dari num, sedangkan d3 dan d4 berisi dua digit terakhir (paling kanan) dari num
	   Contoh: num = 987365, maka d1=9, d2=8, d3=6, dan d4=5 */
	*d4 = num % 10
	*d3 = num % 100 / 10
	if len(num) == 5 {
		*d2 = num / 1000 % 10
		*d1 = num / 10000
	} else {
		*d2 = num / 10000 % 10
		*d1 = num / 100000
	}
}

func valid(ns int) bool {
	/* mengembalikan true apabila nomor seri voucher ns valid, atau false apabila tidak*/
	var length int = len(ns)
	var d1, d2, d3, d4, dT int
	digitAwalAkhir(ns, &d1, &d2, &d3, &d4)
	digitTengah(ns, &dT)
	return (length == 5 || length == 6) && (d1*d2 == d3*d4) && (dT%2 == 0)
}

func digitTengah(num int, dTengah *int) {
	/* I.S. terdefinisi bilangan bulat num yang menyatakan 5 atau 6 digit nomor seri voucher
	   F.S. dTengah berisi digit tengah, artinya selain dua digit pertama dan terakhir,
	   Contoh: num = 12345 maka dTengah = 3, sedangkan num = 345678 maka dTengah = 56*/
	if len(num) == 5 {
		*dTengah = num % 1000 / 100
	} else {
		*dTengah = num / 100 % 100
	}
}
