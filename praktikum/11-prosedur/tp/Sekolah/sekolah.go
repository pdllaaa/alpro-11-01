package main

import "fmt"

const NMAX int = 1001

type tabMurid [NMAX]string

// cariMurid mencari nama pada T[0..n-1].
// F.S. ketemu = true dan posisi berisi urutan absen (mulai dari 1) jika nama ditemukan,
// jika tidak ketemu = false.
func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	var i int
	*posisi = 0
	*ketemu = false
	for i = 0; i < n && !*ketemu; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1
		}
	}
}

// hitungHuruf menghitung banyaknya nama pada T[0..n-1] yang diawali huruf h.
func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	var i int
	*jumlah = 0
	for i = 0; i < n; i++ {
		if len(T[i]) > 0 && T[i][0] == h {
			*jumlah = *jumlah + 1
		}
	}
}

func main() {
	var T tabMurid
	var n, i, posisi, jumlah int
	var nama string
	var ketemu bool

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
	fmt.Scan(&nama)

	cariMurid(T, n, nama, &posisi, &ketemu)
	if ketemu {
		fmt.Println("Murid terdaftar dan berada di urutan absen ke-" + fmt.Sprint(posisi))
	} else {
		fmt.Println("Murid tidak terdaftar")
	}

	hitungHuruf(T, n, nama[0], &jumlah)
	fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", nama[0], jumlah)
}