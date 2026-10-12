package main

import "fmt"

func main() {
	var n int
	var teks string

	fmt.Scan(&n)
	fmt.Scan(&teks)

	for i := 0; i < n; i++ {
		fmt.Println(teks)
	}
}