package main

import "fmt"

const NMAX int = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer [NMAX]player

// rataRata mengembalikan rata-rata skor dari n player
func rataRata(T tabPlayer, n int) float64 {
	var total int = 0
	var i int
	for i = 0; i < n; i++ {
		total = total + T[i].score
	}
	return float64(total) / float64(n)
}

// idxTertinggi mengembalikan indeks player dengan skor tertinggi
// (jika ada skor sama, diambil yang pertama muncul)
func idxTertinggi(T tabPlayer, n int) int {
	var idx int = 0
	var i int
	for i = 1; i < n; i++ {
		if T[i].score > T[idx].score {
			idx = i
		}
	}
	return idx
}

// hitungDiAtas mengembalikan banyaknya player dengan skor lebih dari batas
func hitungDiAtas(T tabPlayer, n int, batas float64) int {
	var jumlah int = 0
	var i int
	for i = 0; i < n; i++ {
		if float64(T[i].score) > batas {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var T tabPlayer
	var n, i int

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i].nickname, &T[i].score)
	}

	var rata float64 = rataRata(T, n)
	var idx int = idxTertinggi(T, n)

	fmt.Printf("Rata-rata skor: %.2f\n", rata)
	fmt.Printf("Player terbaik: %s (%d)\n", T[idx].nickname, T[idx].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", hitungDiAtas(T, n, rata))
}