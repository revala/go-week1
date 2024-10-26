package main

import "fmt"

func main() {
	// Mahasiswa 1
	var mahasiswa1 = "Newt"

	//Mendeklarasikan nilai dan rata-rata mahasiswa 1
	nilai1 := [...]int{90, 89, 78}
	totalNilai1 := 0

	for _, nilai := range nilai1 {
		totalNilai1 = totalNilai1 + nilai
	}

	rataRata1 := float32(totalNilai1 / (len(nilai1)))

	// Nilai Huruf mahasiswa 1
	var nilaiHuruf string
	if rataRata1 >= 80 {
		nilaiHuruf = "A"
	} else if rataRata1 >= 70 {
		nilaiHuruf = "B"
	} else if rataRata1 >= 60 {
		nilaiHuruf = "C"
	} else if rataRata1 >= 50 {
		nilaiHuruf = "D"
	} else if rataRata1 < 50 {
		nilaiHuruf = "E"

	} else {
		fmt.Println("Nilai tidak ada, tolong input lagi")
	}

	//Output
	fmt.Println("Nama mahasiswa:", mahasiswa1, "Rata-rata:", rataRata1, "Nilai Huruf:", nilaiHuruf)

	// Mahasiswa 2
	var mahasiswa2 = "Thomas"

	//Mendeklarasikan nilai dan rata-rata mahasiswa 1
	nilai2 := [...]int{88, 80, 60}
	totalNilai2 := 0

	for _, nilais := range nilai2 {
		totalNilai2 = totalNilai2 + nilais
	}

	rataRata2 := float32(totalNilai2 / (len(nilai2)))

	// Nilai Huruf mahasiswa 1
	var nilaiHuruf2 string
	if rataRata2 >= 80 {
		nilaiHuruf2 = "A"
	} else if rataRata2 >= 70 {
		nilaiHuruf2 = "B"
	} else if rataRata2 >= 60 {
		nilaiHuruf2 = "C"
	} else if rataRata2 >= 50 {
		nilaiHuruf2 = "D"
	} else if rataRata2 < 50 {
		nilaiHuruf2 = "E"

	} else {
		fmt.Println("Nilai tidak ada, tolong input lagi")
	}

	//Output
	fmt.Println("Nama mahasiswa:", mahasiswa2, "Rata-rata:", rataRata2, "Nilai Huruf:", nilaiHuruf2)

	//Nilai tertinggi
	if rataRata1 > rataRata2 {
		fmt.Println("Mahasiswa dengan nilai rata-rata tertinggi:", mahasiswa1, "dengan rata-rata:", rataRata1)
	} else if rataRata1 < rataRata2 {
		fmt.Println("Mahasiswa dengan nilai rata-rata tertinggi:", mahasiswa2, "dengan rata-rata:", rataRata2)
	} else {
		fmt.Println("Datà tidak ditemukan, tolong input lagi")
	}
}
