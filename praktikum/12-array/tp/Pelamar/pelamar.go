package main

import "fmt"

const NMAX int = 100

type applicant struct {
	id, name     string
	testGrade    float64
	testDuration int
}

type applicants [NMAX]applicant

// IS : A[0..n-1] dan batas terdefinisi
// FS : B[0..m-1] berisi pelamar dari A dengan testGrade >= batas,
//      sesuai urutan kemunculan pada A
func salinLolos(A applicants, n int, batas float64, B *applicants, m *int) {
	var i int
	*m = 0
	for i = 0; i < n; i++ {
		if A[i].testGrade >= batas {
			B[*m] = A[i] // salin seluruh struct ke array kedua
			*m = *m + 1
		}
	}
}

func rataDurasi(B applicants, m int) float64 {
	var i, total int
	if m == 0 {
		return 0
	}
	for i = 0; i < m; i++ {
		total = total + B[i].testDuration
	}
	return float64(total) / float64(m)
}

func main() {
	var A, B applicants
	var n, m, i int
	var batas float64
	var id string

	// Baca data pelamar sampai "END"
	n = 0
	fmt.Scan(&id)
	for id != "END" && n < len(A) {
		A[n].id = id
		fmt.Scan(&A[n].name, &A[n].testGrade, &A[n].testDuration)
		n++
		fmt.Scan(&id)
	}
	fmt.Scan(&batas)

	salinLolos(A, n, batas, &B, &m)

	fmt.Println("Data awal:")
	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %.1f %d\n", A[i].id, A[i].name, A[i].testGrade, A[i].testDuration)
	}

	fmt.Printf("Pelamar lolos (nilai >= %.1f):\n", batas)
	if m == 0 {
		fmt.Println("Tidak ada pelamar yang lolos")
	} else {
		for i = 0; i < m; i++ {
			fmt.Printf("%s %s %.1f %d\n", B[i].id, B[i].name, B[i].testGrade, B[i].testDuration)
		}
		fmt.Printf("Jumlah pelamar lolos: %d\n", m)
		fmt.Printf("Rata-rata durasi: %.2f menit\n", rataDurasi(B, m))
	}
}