package main

import "fmt"

func main() {
	var x, digit, sebelumnya int
	var cantik bool

	fmt.Scan(&x)

	sebelumnya = 10
	cantik = true

	for x > 0 {
		digit = x % 10

		if digit > sebelumnya {
			cantik = false
			break
		}

		sebelumnya = digit
		x = x / 10
	}

	fmt.Println(cantik)
}