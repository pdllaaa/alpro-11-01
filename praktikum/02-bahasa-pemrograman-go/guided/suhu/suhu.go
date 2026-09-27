package main

import "fmt"

func main() {
	var celsius float64

	// Membaca masukan
	fmt.Scan(&celsius)

	// Menghitung konversi suhu
	reamur := celsius * 4.0 / 5.0
	fahrenheit := celsius * 9.0 / 5.0 + 32.0
	kelvin := celsius + 273.15

	// Menampilkan output dengan label teks penjelasan
	fmt.Println("Reamur:", reamur)
	fmt.Println("Fahrenheit:", fahrenheit)
	fmt.Println("Kelvin:", kelvin)
}