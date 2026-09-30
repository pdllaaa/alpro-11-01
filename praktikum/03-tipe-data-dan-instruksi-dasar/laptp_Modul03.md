# <h1 align="center">Tugas Pendahuluan Modul 03 - “VARIABEL DAN OPERATOR” Bahasa Pemrograman Go (Golang)</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

### 1. sisa_kue.go

```go
package main

import "fmt"

func main() {
	var y, x int

	fmt.Scan(&y, &x)

	sisa := y % x

	fmt.Println(sisa)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisa/outputsisakue.png)


#### Deskripsi
[Pada program `sisa_kue`, saya memasukkan dua bilangan bulat yang menyatakan jumlah kue dan jumlah anggota keluarga. Setelah itu, kedua nilai dibagi sama rata untuk mendapatkan total jumlah kue yang tersisa setelah dibagi rata. Hasil yang ditampilkan berupa jumlah kue sebelum dibagi dan jumlah kue setelah dibagi ke setiap anggota keluarga.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya mempraktikkan cara menerima input, melakukan perhitungan sederhana, dan menampilkan hasil menggunakan `fmt.Println`]

### 2. bool.go

```go
package main

import "fmt"

func main() {
	var nilai bool
	fmt.Scan(&nilai)

	fmt.Println(nilai)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/bool/outputbool.png)


#### Deskripsi
[Pada program `bool`, saya memasukkan nilai boolean true atau boolean false. Hasil yang ditampilkan berupa nilai boolean true yang sudah dibaca lalu dicetak kembali, dan nilai boolean false yang sudah dibaca lalu dicetak kembali.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya mempraktikkan cara menerima input, melakukan perhitungan sederhana, dan menampilkan hasil menggunakan `fmt.Println`]

### 3. mill.go

```go
package main

import "fmt"

func main() {
	var mill float64

	fmt.Scan(&mill)
	
	km := mill * 1.6

	fmt.Printf("%.1f", km)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/konversi/outputmill.png)


#### Deskripsi
[Pada program `mill`, saya memasukkan nilai dalam mill berupa bilangan desimal menggunakan `float64` lalu mengonversinya ke kilometer dan menampilkan hasil konversi dengan 1 angka di belakang koma (megunakan fmt.Printf dengan format %.1f).

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya mempraktikkan cara menerima input, melakukan perhitungan sederhana, dan menampilkan hasil menggunakan `fmt.Printf`]

## Kesimpulan
[Setelah saya mengerjakan Tugas Pendahuluan Modul 03, saya dapat memahami dasar pemrograman Go, mulai dari penggunaan package main dan func main(), variabel, tipe data, input dengan fmt.Scan, serta output dengan fmt.Println dan fmt.Printf. Saya juga memahami penggunaan operator aritmatika seperti perkalian dan sisa bagi. Selain itu, saya mengetahui bahwa program Go disimpan dengan ekstensi .go. Materi tersebut kemudian diterapkan pada program guided dan unguided untuk menerima input, melakukan proses perhitungan, dan menampilkan hasil.]