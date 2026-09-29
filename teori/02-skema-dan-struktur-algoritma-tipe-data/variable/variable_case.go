package main

import "fmt"

func main() {
	var name string

	name = "Yudha Islami Sulistya"
	fmt.Println("Nama : ", name)

	var lastName string = "Sulistya"
	fmt.Println("Nama Belakang : ", lastName)

	middleName := "Islami"
	fmt.Println("Nama Tengah : ", middleName)

	var (
		fullName = "Yudha Islami Sulistya"
		firstName = "Yudha"
	)
	fmt.Println("Nama Lengkap : ", fullName)
	fmt.Println("Nama Depan : ", firstName)
}