package main

import "fmt"

func main() {
	var sebelumnya, sekarang int
	var tertinggi, terendah, total, n int

	fmt.Scan(&sebelumnya)
	tertinggi = sebelumnya
	terendah = sebelumnya
	total = sebelumnya
	n = 1

	for {
		fmt.Scan(&sekarang)
		if sekarang == sebelumnya {
			break
		}

		if sekarang > tertinggi {
			tertinggi = sekarang
		}
		if sekarang < terendah {
			terendah = sekarang
		}
		total += sekarang
		n++
		sebelumnya = sekarang
	}

	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, float64(total)/float64(n))
}