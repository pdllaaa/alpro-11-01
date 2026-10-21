package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	// Inisialisasi variabel untuk menyimpan data siswa terendah
	var min1Nama, min2Nama string
	// Set nilai awal ke angka yang sangat besar agar mudah tergantikan
	var min1Nilai, min2Nilai int = 999999, 999999

	// Menggunakan perulangan 'while' di Go (for dengan kondisi)
	i := 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		// Cek apakah nilai lebih kecil dari nilai terendah pertama
		if nilai < min1Nilai {
			// Geser nilai terendah pertama lama menjadi terendah kedua
			min2Nilai = min1Nilai
			min2Nama = min1Nama

			// Update nilai terendah pertama baru
			min1Nilai = nilai
			min1Nama = nama
		} else if nilai < min2Nilai {
			// Jika hanya lebih kecil dari terendah kedua
			min2Nilai = nilai
			min2Nama = nama
		}

		i++
	}

	fmt.Println(min1Nama)
	fmt.Println(min2Nama)
}