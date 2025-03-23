package main

import "fmt"

func main() {
	var jam, menit int
	var member bool
	var total float64
	fmt.Scan(&jam, &menit, &member)
	hitungSewa(jam, menit, member, &total)
	fmt.Println(total)
}

func durasi(jam, menit int) int {
	/* Mengembalikan durasi dalam jam, apabila diketahui lama sewa dalam jam dan menit.
	Sewa di atas 1 jam, maka kelebihan menit kurang dari 10 menit tidak dihitung penambahan jam.
	Contoh: 1 jam 9 menit, tetap 1 jam. 1 jam 10 menit, menjadi 2 jam  */
	if jam == 0 && menit != 0 {
		return 1
	} else if jam >= 1 {
		if menit >= 10 {
			return jam + 1
		} else {
			return jam
		}
	} else {
		return 0
	}
}

func potongan(durasi int, tarif int) float64 {
	/* Mengembalikan besarnya potongan apabila diketahui durasi dalam jam dan tarif perjamnya.
	Jika durasi lebih besar dari 3 jam, maka besarnya potongan adalah 10%,
	yaitu durasi * tarif * 0.1. Jika durasi kurang dari 3, maka tidak ada potongan, yaitu 0  */
	if durasi > 3 {
		return float64(durasi*tarif) * 0.1
	} else {
		return 0
	}
}

func tarif(member bool) int {
	/* Mengembalikan tarif sewa sesuai dengan status membernya.
	Jika member, maka tarifnya 3500, sedangkan nonmember 5000  */
	if member {
		return 3500
	} else {
		return 5000
	}
}

func hitungSewa(jam, menit int, member bool, biaya *float64) {
	/* I.S. Terdefinisi durasi sewa dalam jam dan menit, serta status
		     membershipnya
	   Proses: Memanggil fungsi durasi untuk menghitung lama sewa,
			  memanggil fungsi tarif untuk menghitung besarnya tarif, dan
		     memanggil fungsi potongan untuk menghitung besarnya diskon.
		     Besarnya totalBiaya dihitung dengan rumus berikut:
		     totalBiaya = lama sewa x besar tarif - diskon
	   F.S. Biaya berisi total biaya sewa setelah dipotong diskon
	*/
	var lamaSewa int = durasi(jam, menit)
	var besarTarif int = tarif(member)
	var diskon float64 = potongan(lamaSewa, besarTarif)
	*biaya = float64(lamaSewa*besarTarif) - diskon
}
