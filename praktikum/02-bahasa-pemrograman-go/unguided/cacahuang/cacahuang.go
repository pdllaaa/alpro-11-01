package main

import "fmt"

func main() {
	//Variabel untuk menyimpan nominal uang
	var uang int

	// Membaca masukan nominal uang
	fmt.Scan(&uang)

	// Menghitung jumlah pecahan uang 10 ribu
	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	// Menghitung jumlah pecahan uang 5 ribu
	limaRibu := sisa / 5000
	sisa = sisa % 5000

	// Menghitung jumlah pecahan uang 1 ribu
	seribu := sisa / 1000

	// Menampilkan hasil perhitungan
	fmt.Println(sepuluhRibu, limaRibu, seribu)
}
