package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func kombinasi(x, y int) int {
	n := x
	r := y

	if x < y {
		n = y
		r = x
	}

	hasil := faktorial(n) / (faktorial(r) * faktorial(n-r))

	return hasil
}

func main() {
	var x, y int
	
	fmt.Scan(&x, &y)

	hasilX := faktorial(x)
    hasilY := faktorial(y)
    hasilKombinasi := kombinasi(x, y)

	fmt.Println(hasilX, hasilY, hasilKombinasi)
}