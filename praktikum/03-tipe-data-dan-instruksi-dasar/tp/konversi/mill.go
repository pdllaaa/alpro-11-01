package main

import "fmt"

func main() {
	var mill float64

	fmt.Scan(&mill)
	
	km := mill * 1.6

	fmt.Printf("%.1f", km)
}