package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	//
	sum := 0
	i := 1

	for i <= n {
		sum += i //menambahkan nilai i ke dalam total (sum)
		i++ //Menaikkan nilai i sebesar 1
	}

	fmt.Println(sum)
}