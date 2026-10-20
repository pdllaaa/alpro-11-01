package main

import "fmt"

func main() {
	var x, angka int
	var digit, kuadrat, total int

	fmt.Scan(&x)

	angka = x
	total = 0

	for x > 0 {
		digit = x % 10
		kuadrat = digit * digit
		total = total + kuadrat
		x = x / 10
	}

	x = angka

	for x > 0 {
		digit = x % 10
		kuadrat = digit * digit

		fmt.Print(kuadrat)

		x = x / 10

		if x > 0 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
	fmt.Println(total)
}