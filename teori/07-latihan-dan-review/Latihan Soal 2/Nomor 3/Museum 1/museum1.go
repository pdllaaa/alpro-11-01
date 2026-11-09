package main

import "fmt"

func main() {
	var pengunjung, sebelumnya int
	var total, hari, kenaikan int

	for {
		fmt.Scan(&pengunjung)

		if pengunjung < 0 || pengunjung > 200 {
			break
		}

		total += pengunjung
		hari++

		if hari > 1 && pengunjung > sebelumnya {
			kenaikan++
		}

		sebelumnya = pengunjung
	}

	rataRata := 0.0

	if hari > 0 {
		rataRata = float64(total) / float64(hari)
	}

	fmt.Printf("%d %.2f\n", kenaikan, rataRata)
}