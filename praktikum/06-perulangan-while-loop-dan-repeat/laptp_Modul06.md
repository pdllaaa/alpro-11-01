# <h1 align="center">Tugas Pendahuluan Modul 06 - “PERULANGAN WHILE” Bahasa Pemrograman Go (Golang)</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

### 1. lulus.go

```go
package main

import "fmt"

func main() {
	var nilai int

	for {
		fmt.Print("Masukan nilai Anda: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			break
		}

		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}
	}

	fmt.Println("Program selesai.")
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/Lulus/output%20lulus.png)


#### Deskripsi
Pada program `lulus.go`, saya menggunakan variabel `nilai` untuk menyimpan angka ujian yang dimasukkan oleh user. Setelah itu, saya menggunakan perulangan `for` tanpa kondisi yang bertindak seperti *while loop* agar program bisa meminta input secara terus-menerus. Jadi setiap kali nilai dimasukkan, sistem akan langsung mengeceknya, jika nilai mencapai angka **60 atau lebih**, maka siswa **dinyatakan lulus ujian**. Namun, jika nilai berada **di bawah 60**, siswa **dinyatakan tidak lulus ujian**. Perulangan baru akan berhenti beroperasi ketika user memasukkan angka `-1` sebagai sinyal keluar, lalu program akan menampilkan pesan penutup yaitu **Program selesai.**.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya belajar cara menggunakan perulangan interaktif yang terus berjalan sampai ada perintah berhenti (*break*), serta menerapkan logika percabangan bercabang untuk menentukan kelulusan menggunakan bahasa Go.

### 2. deret.go

```go
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
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/Deret/output%20deret.png)


#### Deskripsi
Pada program `deret.go`, saya menggunakan variabel `n` untuk menyimpan bilangan bulat positif yang akan dimasukkan oleh user. Selanjutnya, program menyiapkan variabel `sum` yang bernilai `0` untuk menampung total hitungan dan variabel `i` bernilai `1` sebagai angka awal. Selama nilai `i` belum melewati `n`, program akan terus menambahkan nilai `i` ke dalam `sum` dan menaikkan nilai `i` sebesar 1 di setiap putarannya. Setelah perulangan selesai karena nilai `i` sudah lebih besar dari `n`, program akan langsung mencetak total hasil akhir yang tersimpan di dalam variabel `sum`.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya belajar cara mengimplementasikan logika kontrol perulangan *conditional* menggunakan kata kunci `for` di bahasa Go, karena bahasa ini tidak menyediakan kata kunci `while` secara bawaan.

### 3. duaterendah.go

```go
package main

import (
	"fmt"
)

func main() {
	var n int
	
	fmt.Scan(&n)

	// Inisialisasi variabel untuk menyimpan data siswa terendah
	var min1Nama, min2Nama string
	// Set nilai awal ke angka yang sangat besar agar mudah tergantikan
	var min1Nilai, min2Nilai int = 999999, 999999

	// Menggunakan perulangan 'while' di Go (for dengan kondisi)
	i := 0
	for i < n {
		var nama string
		var nilai int
		fmt.Scan(&nama, &nilai)

		// Cek apakah nilai lebih kecil dari nilai terendah pertama
		if nilai < min1Nilai {
			// Geser nilai terendah pertama lama menjadi terendah kedua
			min2Nilai = min1Nilai
			min2Nama = min1Nama

			// Update nilai terendah pertama baru
			min1Nilai = nilai
			min1Nama = nama
		} else if nilai < min2Nilai {
			min2Nilai = nilai
			min2Nama = nama
		}

		i++
	}

	fmt.Println(min1Nama)
	fmt.Println(min2Nama)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/Dua%20Terendah/output%20duaterendah.png)

#### Deskripsi
Pada program `duaterendah.go`, langkah pertama dimulai dengan menggunakan variabel `n` untuk menyimpan bilangan bulat positif berupa jumlah siswa yang dimasukkan oleh user. Selanjutnya, program menyiapkan variabel `nama` dan variabel `nilai` untuk mencatat data siswa, serta variabel `i` yang bertindak sebagai angka awal. Di dalam perulangan, selama nilai `i` belum melewati `n`, program akan terus membaca input nama dan nilai siswa, melakukan pengecekan untuk memperbarui data dua siswa dengan nilai paling kecil, serta menaikkan `nilai i sebesar 1` di setiap putarannya. Setelah perulangan selesai atau ketika nilai `i` sudah tidak lebih kecil dari `n`, program akan berhenti dan langsung mencetak `nama-nama siswa` dengan `nilai terendah` yang tersimpan di dalam sistem.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya belajar cara mengimplementasikan logika kontrol perulangan `for` di bahasa Go serta cara melakukan penyaringan data secara dinamis di dalam perulangan tersebut.

### 4. biaya.go

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	biaya := 10000000.0

	if n > 1 {
		biaya = biaya * math.Pow(1.06, float64(n-1))
	}

	fmt.Printf("%.0f\n", biaya)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/Biaya/output%20biaya.png)


#### Deskripsi
Pada program `biaya.go`, saya menggunakan variabel `n` dengan tipe data int untuk menyimpan bilangan bulat positif yang menyatakan tahun ke-n yang dimasukkan oleh user lewat `fmt.Scan`. Selanjutnya, program menginisialisasi variabel `biaya` dengan tipe data `float64` yang bernilai `10000000.0` sebagai representasi biaya kuliah awal pada tahun pertama.

Untuk menghitung kenaikan biaya kuliah sebesar 6% setiap tahunnya, program memanfaatkan fungsi esponensial `math.Pow(1.06, float649n-1))`. Logika ini memastikan bahwa perkiraan biaya kumulatif dihitung secara akurat berdasarkan selisih tahun dari tahun pertama (n-1). Setelah proses perhitungan selesai, program langsung mencetak hasil akhirnya menggunakan format fungsi `fmt.Printf("%.0f\n", biaya)` untuk membulatkan nilai desimal float tersebut ke mata uang rupiah bulat terdekat tanpa angka di belakang koma.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya belajar cara mengimplementasikan operasi perpangkatan menggunakan library `math` di bahasa Go, melakukan konversi tipe data secara aman dari `int` ke `float64`, serta menformat output desimal agar sesuai dengan spesifikasi pembulatan yang diminta oleh soal.

## Kesimpulan
Setelah mengerjakan Tugas Pendahuluan Modul 06, saya lebih memahami penggunaan perulangan for sebagai pengganti while dalam bahasa Go. Saya belajar menggunakan perulangan untuk mengecek kelulusan, menghitung jumlah deret, mencari dua nilai terendah, dan menghitung kenaikan biaya kuliah. Dari tugas ini, saya jadi lebih paham cara menggunakan perulangan untuk menyelesaikan masalah dalam program.