package main

import "fmt"

func hitung(x, y int) int { // x dan y adalah sebuah parameter
	hasil := 2*x + y - 5
	return hasil
}

func kurang(x, y float32) float32 {
	hasil := x - y
	return hasil
}

func ambangAtas() int {
	return 60
}

// hitung itu adalah nama fungsi
// {hasil/proses} ini adalah instruksi atau perintah atau proses
// kamus lokal variabel x, y dengan tipe integer

func lulus(nilai int) bool {
	return nilai >= ambangAtas()
}

func statusLulus(nilai int) string {
	if lulus(nilai) {
		return "Lulus"
	} else {
		return "Tidak Lulus"
	}
}

func main() {
	a := 10
	fmt.Println(lulus(50))
	fmt.Println(statusLulus(80))
	fmt.Println(ambangAtas())
	fmt.Println(hitung(a, 5) + 20) // a dan 5 ini adalah sebuah argumen
	fmt.Println(hitung(20, 10))
	fmt.Println(hitung(30, 1))
}