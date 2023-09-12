// Buat sebuah struct dengan nama Student yang mempunyai properti name dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa program menunjukkan skor rata-rata, skor minimum dan maksimal dari 5 siswa
package main

import "fmt"

type Student struct {
	name  string
	score int
}

func main() {
	var paraSiswa = []Student{
		{name: "Rizki", score: 80},
		{name: "Kobar", score: 75},
		{name: "Ismail", score: 70},
		{name: "Umam", score: 75},
		{name: "Yopan", score: 60},
	}
	// variabel persiapan
	totalScore := 0
	totalSiswa := len(paraSiswa)
	nilaiTerendah := 100
	siswaNilaiTerendah := ""
	nilaiTertinggi := 0
	siswaNilaiTertinggi := ""

	for _, siswa := range paraSiswa {
		totalScore += siswa.score
		skor := siswa.score

		// mencari siswa dengan nilai terendah
		if skor < nilaiTerendah {
			nilaiTerendah = skor
			siswaNilaiTerendah = siswa.name
		}

		// mencari siswa dengan nilai tertinggi
		if skor > nilaiTertinggi {
			nilaiTertinggi = skor
			siswaNilaiTertinggi = siswa.name
		}
	}
	// mencari nilai rata-rata
	nilaiRata := totalScore / totalSiswa

	fmt.Println("nilai rata-rata :", nilaiRata)
	fmt.Println("nilai terendah :", siswaNilaiTerendah, nilaiTerendah)
	fmt.Println("nilai tertinggi :", siswaNilaiTertinggi, nilaiTertinggi)
}
