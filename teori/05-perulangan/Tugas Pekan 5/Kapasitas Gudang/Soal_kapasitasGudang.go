package main

import "fmt"

func main() {
	var k, v, total int

	fmt.Scan(&k)

	total = 0

	for i := 0; i < 5; i++ {
		fmt.Scan(&v)

		total += v

		if total >= k {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
}