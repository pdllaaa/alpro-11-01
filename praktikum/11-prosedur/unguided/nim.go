package main

import "fmt"

const NMAX int = 2002

type tabNIM [NMAX]int

// cariNIM mencari NIM x pada T[0..n-1]
// IS: T[0..n-1] dan x terdefinisi
// FS: idx berisi indeks NIM x pada T, atau -1 jika tidak ditemukan
func cariNIM(T tabNIM, n, x int, idx *int) {
	var i int
	*idx = -1
	for i = 0; i < n && *idx == -1; i++ {
		if T[i] == x {
			*idx = i
		}
	}
}

// ubahNIM mengganti isi T[idx] dengan NIM baru
// IS: T terdefinisi dan idx adalah indeks yang valid
// FS: T[idx] berisi NIM baru
func ubahNIM(T *tabNIM, idx, baru int) {
	T[idx] = baru
}

// cetak menampilkan setiap NIM pada T, satu per baris
// IS: T[0..n-1] terdefinisi
// FS: setiap NIM pada T tercetak satu per baris
func cetak(T tabNIM, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var T tabNIM
	var n, i, x, baru, idx int

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
	fmt.Scan(&x)

	cariNIM(T, n, x, &idx)

	if idx != -1 {
		fmt.Scan(&baru)
		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", x, idx+1)
		ubahNIM(&T, idx, baru)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}