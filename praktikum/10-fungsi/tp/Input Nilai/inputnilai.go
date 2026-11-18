package main

import "fmt"

const NMAX int = 100

type student struct {
	nim, name string
	grade     float64
}

type students [NMAX]student

// cariNIM mengembalikan indeks data dengan NIM tersebut, atau -1 jika tidak ada
func cariNIM(T students, n int, nim string) int {
	var i int
	for i = 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}

// rataRata mengembalikan rata-rata nilai dari n data pertama pada T
func rataRata(T students, n int) float64 {
	var i int
	var total float64
	if n == 0 {
		return 0
	}
	for i = 0; i < n; i++ {
		total = total + T[i].grade
	}
	return total / float64(n)
}

func main() {
	var T students
	var n, i int
	var nim, name string
	var grade float64

	fmt.Scan(&nim)
	for nim != "STOP" && n < NMAX {
		fmt.Scan(&name, &grade)
		// hanya simpan data jika NIM belum pernah tersimpan
		if cariNIM(T, n, nim) == -1 {
			T[n].nim = nim
			T[n].name = name
			T[n].grade = grade
			n++
		}
		fmt.Scan(&nim)
	}

	for i = 0; i < n; i++ {
		fmt.Println(T[i].nim, T[i].name, T[i].grade)
	}
	fmt.Println()
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}