package main

import "fmt"

// Tipe bentukan struct mahasiswa dengan field nama, NIM, dan nilai

type mahasiswa struct {
	nama, NIM string
	nilai     float64
}

func main() {
	var mhs1, mhs2, mhs3 mahasiswa
	baca(&mhs1, &mhs2, &mhs3)
	cetak(mhs1, mhs2, mhs3)
	fmt.Println("Rata-rata nilai:", rata_rata_nilai(mhs1, mhs2, mhs3))
	mhs_max_nilai(mhs1, mhs2, mhs3)
}

func baca(m1, m2, m3 *mahasiswa) {
	/* IS: m1, m2, m3 terdefinisi sembarang
	   Proses: membaca dari piranti masukan data nama, NIM, dan nilai dari m1, m2, dan m3
	   FS: m1, m2, m3 berisi nilai
	*/
	fmt.Scan(&m1.nama, &m1.NIM, &m1.nilai)
	fmt.Scan(&m2.nama, &m2.NIM, &m2.nilai)
	fmt.Scan(&m3.nama, &m3.NIM, &m3.nilai)
}

func cetak(m1, m2, m3 mahasiswa) {
	/* IS: m1, m2, m3 terdefinisi
	   FS: Semua field dari m1, m2, m3 tercetak di layar
	*/
	fmt.Println(m1.nama, m1.NIM, m1.nilai)
	fmt.Println(m2.nama, m2.NIM, m2.nilai)
	fmt.Println(m3.nama, m3.NIM, m3.nilai)
}

func rata_rata_nilai(m1, m2, m3 mahasiswa) float64 {
	/* Mengembalikan rata-rata nilai dari 3 mahasiswa m1, m2, dan m3 */
	return (m1.nilai + m2.nilai + m3.nilai) / 3
}

func mhs_max_nilai(m1, m2, m3 mahasiswa) {
	/* IS: m1, m2, dan m3 terdefinisi
	   FS: Tercetak di layar nama mahasiswa dengan nilai tertinggi dengan format
	       "Mahasiswa dengan nilai tertinggi: <nama mahasiswa> <NIM mahasiswa>"
	       Asumsi nilai mahasiswa bersifat unik, sehingga hanya ada 1 nilai tertinggi
	*/
	var max mahasiswa
	max = m1
	if m2.nilai > max.nilai {
		max = m2
	}
	if m3.nilai > max.nilai {
		max = m3
	}
	fmt.Printf("Mahasiswa dengan nilai tertinggi %v: %s %s\n", max.nilai, max.nama, max.NIM)
}
