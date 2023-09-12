// Buat sebuah struct dengan nama Student yang mempunyai properti name dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa program menunjukkan skor rata-rata, skor minimum dan maksimal dari 5 siswa
package main

import "fmt"

type Student struct {
	Name  []string
	Score []int
}

func (paraSiswa Student) NilaiRata() string {
	var sum int
	for _, v := range paraSiswa.Score {
		sum += v
	}
	result := float32(sum) / float32(len(paraSiswa.Score))
	return fmt.Sprintf("nilai rata-rata siswa : %v", result)
}

func (paraSiswa Student) Minimum() string {
	nilaiTerendah := paraSiswa.Score[0]
	indexMinimum := 0
	for i, v := range paraSiswa.Score {
		if v < nilaiTerendah {
			nilaiTerendah = v
			indexMinimum = i
		}
	}
	return fmt.Sprintf("nilai terendah siswa : %v (%v)", paraSiswa.Name[indexMinimum], nilaiTerendah)
}
func (paraSiswa Student) Maksimum() string {
	nilaiTertinggi := paraSiswa.Score[0]
	indexMaksimum := 0
	for i, v := range paraSiswa.Score {
		if v < nilaiTertinggi {
			nilaiTertinggi = v
			indexMaksimum = i
		}
	}
	return fmt.Sprintf("nilai tertinggi siswa : %v (%v)", paraSiswa.Name[indexMaksimum], nilaiTertinggi)
}

func main() {
	fmt.Println("Input :")
	var paraSiswa Student

	for i := 1; i <= 5; i++ {
		fmt.Printf("input %v nama siswa : ", i)
		paraSiswa.Name = append(paraSiswa.Name, "")
		fmt.Scanln(&paraSiswa.Name[i-1])
		fmt.Printf("input %v nilai siswa : ", i)
		paraSiswa.Score = append(paraSiswa.Score, 0)
		fmt.Scanln(&paraSiswa.Score[i-1])
	}
	// Inputnya
	// Rizki 80
	// Kobar 75
	// Ismail 70
	// Umam 75
	// Yopan 60
	fmt.Println(paraSiswa)
	fmt.Println("Output :")
	fmt.Println(paraSiswa.NilaiRata())
	fmt.Println(paraSiswa.Minimum())
	fmt.Println(paraSiswa.Maksimum())
}
