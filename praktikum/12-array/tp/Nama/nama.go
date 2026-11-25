package main

import "fmt"

const NMAX int = 9999

type arrString [NMAX]string

// IS : n terdefinisi
// FS : A[0..n-1] terisi dari masukan
func bacaData(A *arrString, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

// FS : seluruh elemen tampil dipisahkan koma dan diakhiri titik
func cetakData(A arrString, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print(A[i])
		if i < n-1 {
			fmt.Print(", ")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

// IS : A[0..n-1] terdefinisi, 1 <= p <= n+1
// FS : x berada pada posisi ke-p, elemen setelahnya bergeser ke kanan,
//      n bertambah 1
func sisipData(A *arrString, n *int, p int, x string) {
	var i int
	// geser elemen dari belakang agar tidak saling menimpa
	for i = *n; i > p-1; i-- {
		A[i] = A[i-1]
	}
	A[p-1] = x
	*n = *n + 1
}

// IS : A[0..n-1] terdefinisi, 1 <= q <= n
// FS : elemen pada posisi ke-q terhapus, elemen setelahnya bergeser
//      ke kiri, n berkurang 1
func hapusData(A *arrString, n *int, q int) {
	var i int
	for i = q - 1; i < *n-1; i++ {
		A[i] = A[i+1]
	}
	*n = *n - 1
}

func main() {
	var A arrString
	var n, p, q int
	var x string

	fmt.Scan(&n)
	bacaData(&A, n)
	fmt.Scan(&p, &x)
	fmt.Scan(&q)

	fmt.Print("Data awal: ")
	cetakData(A, n)

	sisipData(&A, &n, p, x)
	fmt.Print("Setelah penyisipan: ")
	cetakData(A, n)

	hapusData(&A, &n, q)
	fmt.Print("Setelah penghapusan: ")
	cetakData(A, n)
}