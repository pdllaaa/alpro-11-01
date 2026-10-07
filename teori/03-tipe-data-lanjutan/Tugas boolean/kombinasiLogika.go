package main

import "fmt"

func main() {
	var p, q int

	fmt.Scan(&p, &q)

	fmt.Println(p%2 == 0 || q%2 == 0)
	fmt.Println(p%2 != 0 && q%2 != 0)
	fmt.Println(!(p == q))
}