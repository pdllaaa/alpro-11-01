package main

import "fmt"

func main() {
	var n, i, bilangan, jumlah, d1, d4 int
	fmt.Scan(&n)
	jumlah = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&bilangan)
		d1 = bilangan / 1000
		d4 = bilangan % 10
		jumlah += d1 + d4
	}
	fmt.Printf("%d ", jumlah)
}

//jawaban untuk input 3  2022 4321 5578 adalah 22
// d1 = 2022 / 1000 = 2
// d4 = 2022 % 10 = 2
// jumlah = 2 + 2 = 4

// d1 = 4321 / 1000 = 4
// d4 = 4321 % 10 = 1
// jumlah = 4 + 1 + 4 = 9

// d1 = 5578 / 1000 = 5
// d4 = 5578 % 10 = 8
// jumlah = 5 + 8 + 9 = 22
