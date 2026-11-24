package main

import "fmt"

const NMAX int = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa [NMAX]gempa

// kategori mengembalikan kategori gempa berdasarkan magnitudo m
func kategori(m float64) string {
	if m < 4.0 {
		return "Minor"
	} else if m < 6.0 {
		return "Sedang"
	} else if m < 7.0 {
		return "Kuat"
	}
	return "Besar"
}

// rataMagnitudo mengembalikan rata-rata magnitudo dari n gempa
func rataMagnitudo(T tabGempa, n int) float64 {
	var total float64 = 0
	var i int
	for i = 0; i < n; i++ {
		total = total + T[i].magnitudo
	}
	return total / float64(n)
}

// idxTerkuat mengembalikan indeks gempa dengan magnitudo terbesar
// (jika sama, diambil yang pertama muncul)
func idxTerkuat(T tabGempa, n int) int {
	var idx int = 0
	var i int
	for i = 1; i < n; i++ {
		if T[i].magnitudo > T[idx].magnitudo {
			idx = i
		}
	}
	return idx
}

// hitungKategori mengembalikan banyaknya gempa berkategori k
// (memanggil fungsi kategori di dalamnya)
func hitungKategori(T tabGempa, n int, k string) int {
	var jumlah int = 0
	var i int
	for i = 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var T tabGempa
	var n, i int

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}

	for i = 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n", T[i].nama, T[i].tahun, T[i].magnitudo, kategori(T[i].magnitudo))
	}

	fmt.Printf("Rata-rata magnitudo: %.2f\n", rataMagnitudo(T, n))

	var idx int = idxTerkuat(T, n)
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", T[idx].nama, T[idx].tahun, T[idx].magnitudo)
	fmt.Printf("Jumlah gempa kategori Besar: %d\n", hitungKategori(T, n, "Besar"))
}