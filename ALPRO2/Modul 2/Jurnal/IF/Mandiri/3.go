//COMPETITITON JUDGEMENT

//IMPORTANT:
//Disarankan Bentuk Pengerjaan Soal Ini Ada Clue Mirip Answer Box Preload Kalau di LMS jadi Mahasiswa Melengkapi Template Code
//Terutama Bagian Printing

package main

import "fmt"

func main() {

	//Struct Person
	type Person struct {
		firstName string
	}

	//Struct Teams
	type Teams struct {
		competition string
		member1     Person
		member2     Person
		member3     Person
		score       float64
		passed      bool
	}

	//Struct Judges
	type Judges struct {
		initial    string
		scoreGiven float64
	}

	//Deklarasi Variabel
	var t Teams
	var j1, j2, j3 Judges
	var count int = 1
	var compId, teamId string

	//Penginputan Data
	fmt.Printf("Masukkan data dari kelompok ke-%d\n", count)
	fmt.Print("Anggota pertama: ")
	fmt.Scan(&t.member1.firstName)

	//Penginputan Data Dengan "#" Sebagai Sentinel
	for t.member1.firstName != "#" {

		fmt.Print("Anggota kedua: ")
		fmt.Scan(&t.member2.firstName)
		fmt.Print("Anggota ketiga: ")
		fmt.Scan(&t.member3.firstName)
		fmt.Print("Cabang lomba: ")
		fmt.Scan(&t.competition)

		//Konversi Cabang Lomba
		if t.competition == "CP" {
			compId = "111"
		} else if t.competition == "CTF" {
			compId = "112"
		} else if t.competition == "ITB" {
			compId = "113"
		}

		//Pembuatan ID Kelompok
		teamId = compId + t.member1.firstName + t.member2.firstName + t.member3.firstName
		fmt.Printf("Masukkan data penilaian dari kelompok %s\n", teamId)

		//Data Juri Pertama
		fmt.Println("Masukkan data juri pertama")
		fmt.Print("Inisial: ")
		fmt.Scan(&j1.initial)
		fmt.Print("Skor yang diberikan: ")
		fmt.Scan(&j1.scoreGiven)

		//Data Juri Pertama
		fmt.Println("Masukkan data juri kedua")
		fmt.Print("Inisial: ")
		fmt.Scan(&j2.initial)
		fmt.Print("Skor yang diberikan: ")
		fmt.Scan(&j2.scoreGiven)

		//Data Juri Pertama
		fmt.Println("Masukkan data juri ketiga")
		fmt.Print("Inisial: ")
		fmt.Scan(&j3.initial)
		fmt.Print("Skor yang diberikan: ")
		fmt.Scan(&j3.scoreGiven)

		//Rata-Rata Score
		t.score = (j1.scoreGiven + j2.scoreGiven + j3.scoreGiven) / 3

		//Pengecekan Kelulusan Kelompok
		if t.score >= 80 {
			t.passed = true
		} else {
			t.passed = false
		}

		//Algoritma Output
		fmt.Printf("\nDATA TIM %s\n", teamId)
		fmt.Printf("Anggota:\n%s \n%s \n%s\n", t.member1.firstName, t.member2.firstName, t.member3.firstName)
		fmt.Printf("Dinilai oleh: %s %s %s\n", j1.initial, j2.initial, j3.initial)
		if t.passed {
			fmt.Printf("Lolos dengan nilai: %.2f\n", t.score)
		} else {
			fmt.Printf("Tidak lolos dengan nilai: %.2f\n", t.score)
		}

		//Teknis Loop
		count++
		fmt.Printf("\nMasukkan data dari kelompok ke-%d\n", count)
		fmt.Print("Anggota pertama: ")
		fmt.Scan(&t.member1.firstName)
	}

}
