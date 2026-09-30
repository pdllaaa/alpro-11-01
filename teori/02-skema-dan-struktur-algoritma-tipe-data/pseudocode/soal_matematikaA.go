package main

import "fmt"

func main() {
	var x, y int
	var hasil float64

	fmt.Scan(&x, &y)

	hasil = (1.0 / float64(3*x*x+10)) + float64(10*y) + 7

	fmt.Println(hasil)
}