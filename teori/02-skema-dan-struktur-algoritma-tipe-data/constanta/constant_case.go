package main

import "fmt"

func main() {
	const pi = 3.14

	fmt.Println("Nilai Pi tidak berubah adalah : ", pi)

	//kode program constant 
	const (
		fullName = "Yudha Islami Sulistya"
		firstName = "Yudha"
		lastName  = "Sulistya"
	)
	fmt.Println("Nama Lengkap : ", fullName)
	fmt.Println("Nama Depan : ", firstName)
	fmt.Println("Nama Belakang : ", lastName)
}