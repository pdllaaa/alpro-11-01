package main

import "fmt"

func main() {
	var kendaraan string
	var M, K, T int

	fmt.Scan(&kendaraan)
	for kendaraan == "M" || kendaraan == "K" || kendaraan == "T" {
		if kendaraan == "M" {
			M++
		} else if kendaraan == "K" {
			K++
		} else {
			T++
		}
		fmt.Scan(&kendaraan)
	}

	fmt.Println("Tipe M =", M)
	fmt.Println("Tipe K =", K)
	fmt.Println("Tipe T =", T)
}