package main

import "fmt"

func main() {
	var n, m, x, y int
	var jumlah int

	fmt.Scan(&n, &m, &x, &y)

	for n >= x && m >= y {
		n = n - x
		m = m - y
		jumlah = jumlah + 1
	}

	fmt.Println(jumlah)
}