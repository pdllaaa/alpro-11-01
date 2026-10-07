package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var sepuluhRibu int = x / 10000
	var sisa int = x % 10000

	var limaRibu int = sisa / 5000
	sisa = sisa % 5000

	var seribu int = sisa / 1000

	fmt.Println(sepuluhRibu, limaRibu, seribu)
}
