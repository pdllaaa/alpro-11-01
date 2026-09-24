package main

import "fmt"

func main() {
	// Variabel untuk menyimpan dua bilangan
	var a, b int

	// Membaca masukan dua bilangan
	fmt.Scan(&a, &b)

	// Melakukan operasi aritmatika
	penjumlahan := a + b
	pengurangan := a - b
	perkalian := a * b
	pembagian := a / b
	sisaBagi := a % b

	// Menampilkan hasil operasi aritmatika
	fmt.Println(penjumlahan)
	fmt.Println(pengurangan)
	fmt.Println(perkalian)
	fmt.Println(pembagian)
	fmt.Println(sisaBagi)
}
