package main

import "fmt"

func main() {
	var n, i, lebar, tersempit int
	fmt.Scan(&n)
	fmt.Scan(&tersempit)
	
	for i = 2; i <= n; i++ {
		fmt.Scan(&lebar)
		if lebar < tersempit {
			tersempit = lebar
		}
	}

	fmt.Println(tersempit)
}