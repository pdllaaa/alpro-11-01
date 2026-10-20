
package main

import "fmt"

func main() {
	var m int
	var p, v, ga, fs, ci, tc int
	var total int
	var jumlahLulus int
	var rataRata float64

	fmt.Scan(&m)

	for i := 1; i <= m; i++ {
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		total = p + v + ga + fs + ci + tc
		rataRata = float64(total) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Ya")
			jumlahLulus = jumlahLulus + 1
		} else {
			fmt.Println("Tidak")
		}
	}

	fmt.Println(jumlahLulus)
}