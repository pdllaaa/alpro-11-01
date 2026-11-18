package main

import "fmt"

const NMAX int = 999

type rapot struct {
	matkul string
	nilai  int
}

type tabRapot [NMAX]rapot

// nilaiHuruf mengubah nilai angka menjadi nilai huruf
func nilaiHuruf(x int) string {
	var huruf string
	if x > 80 {
		huruf = "A"
	} else if x > 70 {
		huruf = "AB"
	} else if x > 65 {
		huruf = "B"
	} else if x > 60 {
		huruf = "BC"
	} else if x > 50 {
		huruf = "C"
	} else if x > 40 {
		huruf = "D"
	} else {
		huruf = "E"
	}
	return huruf
}

// rataRata menghitung rata-rata nilai dari n data pertama
func rataRata(T tabRapot, n int) float64 {
	var i, total int
	var hasil float64
	for i = 0; i < n; i++ {
		total = total + T[i].nilai
	}
	if n > 0 {
		hasil = float64(total) / float64(n)
	}
	return hasil
}

func main() {
	var T tabRapot
	var n, i int

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i].matkul, &T[i].nilai)
	}

	for i = 0; i < n; i++ {
		fmt.Println(T[i].matkul, T[i].nilai, nilaiHuruf(T[i].nilai))
	}
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}