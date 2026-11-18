package main

import "fmt"

const NMAX int = 1000

type penduduk struct {
	nama  string
	tahun int
	kota  string
}

type tabPenduduk [NMAX]penduduk

// cariTahun mencari indeks data pertama dengan tahun lahir th, atau -1 jika tidak ada.
func cariTahun(T tabPenduduk, n, th int, idx *int) {
	var i int
	*idx = -1
	for i = 0; i < n && *idx == -1; i++ {
		if T[i].tahun == th {
			*idx = i
		}
	}
}

// tertuaTermuda mencari penduduk dengan tahun lahir terkecil (tua)
// dan tahun lahir terbesar (muda). Syarat: n >= 1.
func tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk) {
	var i int
	*tua = T[0]
	*muda = T[0]
	for i = 1; i < n; i++ {
		if T[i].tahun < (*tua).tahun {
			*tua = T[i]
		}
		if T[i].tahun > (*muda).tahun {
			*muda = T[i]
		}
	}
}

func main() {
	var T tabPenduduk
	var n, i, th, idx int
	var tua, muda penduduk

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].kota)
	}
	fmt.Scan(&th)

	cariTahun(T, n, th, &idx)
	if idx != -1 {
		fmt.Println(T[idx].nama, T[idx].tahun, T[idx].kota, "ditemukan di index ke-"+fmt.Sprint(idx))
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	tertuaTermuda(T, n, &tua, &muda)
	fmt.Printf("Penduduk tertua: %s (%d)\n", tua.nama, tua.tahun)
	fmt.Printf("Penduduk termuda: %s (%d)\n", muda.nama, muda.tahun)
}