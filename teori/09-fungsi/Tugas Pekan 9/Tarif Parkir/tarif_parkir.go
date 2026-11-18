package main

import "fmt"

func tarifPakir(jam int) int {
	if jam <= 2 {
		return 5000
	}

	total := 5000 + (jam-2)*3000

	return total
}

func main() {
	var jam int

	fmt.Scan(&jam)

	fmt.Println(tarifPakir(jam))
}