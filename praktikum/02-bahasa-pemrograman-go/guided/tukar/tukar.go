package main

import "fmt"

func main() {
	var a, b int

	//Membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar nilai a dan b
	a, b = b, a

	//Menampilkan output
	fmt.Println("Nilai a setelah ditukar:", a)
	fmt.Println("Nilai b setelah ditukar:", b)
}
