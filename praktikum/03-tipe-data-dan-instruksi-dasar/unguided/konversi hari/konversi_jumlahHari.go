package main

import "fmt"

func main() {
	var totalHari int

	fmt.Scan(&totalHari)

	// 1 tahun = 12 bulan x 30 hari = 360 hari
	tahun := totalHari / 360
	sisaHari := totalHari % 360

	//1 bulan = 30 hari
	bulan := sisaHari / 30
	sisaHari = sisaHari % 30

	//1 minggu = 7 hari
	minggu := sisaHari / 7
	sisaHari = sisaHari % 7

	//Output empat baris bilangan bulat
	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(minggu)
	fmt.Println(sisaHari)
}