package main

import "fmt"

func main() {
	var L, R int
	fmt.Scan(&L, &R)

	hasil := 0
	for i := L; i <= R; i++ {
		angka := i
		terbalik := 0
		for angka > 0 {
			terbalik = terbalik*10 + angka%10
			angka /= 10
		}
		if terbalik == i {
			hasil++
		}
	}
	fmt.Println(hasil)
}