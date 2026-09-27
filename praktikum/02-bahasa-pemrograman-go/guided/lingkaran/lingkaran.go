package main

import "fmt"

func main() {
	var jariJari float64

	//Membaca input
	fmt.Scan(&jariJari)

	//Menghitung luas lingkaran
	luas := 3.14 * jariJari * jariJari

	//Menampilkan output
	fmt.Println("Luas Lingkaran:", luas)
}