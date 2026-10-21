package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	biaya := 10000000.0

	if n > 1 {
		biaya = biaya * math.Pow(1.06, float64(n-1))
	}

	fmt.Printf("%.0f\n", biaya)
}