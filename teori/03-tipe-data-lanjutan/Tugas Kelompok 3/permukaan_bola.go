package main

import "fmt"

func main() {
	var r float64
	fmt.Scan(&r)
	
	luas := 4 * (22.0 / 7.0) * r * r
	fmt.Println(luas)
}