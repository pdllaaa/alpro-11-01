package main

import "fmt"

func main() {
	var berat, tinggi, bmi float64
	var kategori string

	fmt.Scan(&berat, &tinggi)

	tinggi = tinggi / 100

	bmi = berat / (tinggi * tinggi)

	if bmi < 18.5 {
		kategori = "Kurus"
	} else if bmi < 25 {
		kategori = "Normal"
	} else if bmi < 30 {
		kategori = "Gemuk"
	} else {
		kategori = "Obesitas"
	}

	fmt.Printf("%.2f\n", bmi)
	fmt.Println(kategori)
}