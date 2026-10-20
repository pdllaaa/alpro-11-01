# <h1 align="center">Laporan Praktikum Modul 05 - Perulangan For Bahasa Pemrograman Go (Golang) 
</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

## Dasar Teori

### A. Pengertian Perulangan (Looping)
Perulangan atau looping adalah struktur dalam pemrograman yang  digunakan untuk menjalankan beberapa instruksi secara berulang selama kondisi tertentu masih terpenuhi. Dengan perulangan ini, tidak perlu menulis perintah yang sama secara berkali-kali. Perulangan dapat juga digunakan, misalnya, untuk menampilkan urutan angka, menghitung jumlah bilangan, dan membuat pola karakter.

### B. Perulangan for dalam Bahasa Go
Bahasa Go menggunakan kata kunci for untuk membuat perulangan. Kata kunci ini dapat digunakan dalam beberapa bentuk, sesuai dengan kebutuhan program. Salah satu bentuknya adalah for dengan tiga komponen, yaitu inisialisasi, kondisi, dan pernyataan update. Ketiga bagian tersebut digunakan untuk menentukan nilai awal, memeriksa kondisi, dan mengubah nilai variabel selama perulangan berlangsung.

#### C. Bentuk-Bentuk Perulangan for
Dalam bahasa Go, terdapat empat bentuk penulisan perulangan for. Pertama, for dengan tiga komponen yang digunakan untuk mengatur jalannya perulangan melalui inisialisasi, kondisi, dan update. Kedua, for dengan kondisi saja, yang menjalankan kode selama kondisi masih terpenuhi. Ketiga, for tanpa kondisi, yang menghasilkan perulangan tanpa henti jika tidak dihentikan. Keempat, for-range, yang digunakan untuk mengunjungi seluruh elemen dalam struktur data seperti array, slice, string, atau map.

#### D. Penggunaan break pada Perulangan
Pada perulangan for tanpa kondisi, program dapat berjalan terus-menerus. Untuk menghentikannya ketika kondisi tertentu terpenuhi, dapat digunakan pernyataan break. Dengan demikian, perulangan dapat berhenti sesuai kondisi yang telah ditentukan dalam program.

#### E. Perulangan Bersarang (Nested For-Loop)
Perulangan bersarang (nested for-loop) adalah perulangan for yang berada di dalam perulangan for lainnya. Perulangan ini dapat digunakan untuk membuat pola, matriks, atau tabel. Pada contoh dalam modul, perulangan bersarang digunakan untuk menampilkan pola bintang yang tersusun dalam beberapa baris.

## Guided

### 1. [tabungan.go]

```go
package main

import "fmt"

func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)

	total := 0;
	tabunganHariIni := a;

	for i := 1; i <= n; i++ {
		total = total + tabunganHariIni
		tabunganHariIni = tabunganHariIni + b
	}
	fmt.Println(total)
}	
```
#### Deskripsi
Pada program `tabungan.go`, saya membuat program untuk menghitung total tabungan Budi selama beberapa hari. Program membaca jumlah tabungan awal, kenaikan tabungan setiap hari, dan jumlah hari menabung. Dengan menggunakan perulangan `for`, program menjumlahkan tabungan setiap hari hingga mendapatkan total tabungan Budi.

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_skor](/praktikum/05-perulangan-for-loop/guided/guided-1/output%20tabungan.png)

### 2. [terbesar.go]

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var angka, max int
	fmt.Scan(&angka)
	max = angka

	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)
		if angka > max {
			max = angka
		}
	}
	fmt.Println(max)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_tukar](/praktikum/05-perulangan-for-loop/guided/guided-2/output%20terbesar.png)

#### Deskripsi
Pada program `terbesar.go`, saya membuat program untuk mencari bilangan terbesar dari beberapa bilangan yang dimasukkan oleh user. Program menggunakan perulangan `for` untuk membaca dan membandingkan setiap bilangan. Jika ditemukan bilangan yang lebih besar, nilai terbesar akan diperbarui. Setelah semua bilangan diperiksa, program menampilkan bilangan terbesar.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. [bioskop.go]

```go
package main

import "fmt"

func main() {
	var transaksi, jumlahTiket int
	var hari, jenisFilm string
	var diskon int

	totalPendapatan := 0

	fmt.Scanln(&transaksi)
	for i := 1; i <= transaksi; i++ {
		fmt.Printf("Transaksi ke- %d: ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0;

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if jenisFilm == "Reguler" {
				hargaSatuan = 40000
			} else if jenisFilm == "3D" {
				hargaSatuan = 60000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 80000
			}
		} else if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenisFilm == "Reguler" {
				hargaSatuan = 50000
			} else if jenisFilm == "3D" {
				hargaSatuan = 70000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 90000
			}
		}

		totalHarga := hargaSatuan * jumlahTiket

		fmt.Printf("Harga Sebelum Diskon: %d\n", totalHarga)

		if jumlahTiket >= 5 {
			diskon = (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}
		fmt.Printf("Diskon sebesar: %d\n", diskon)
		fmt.Printf("Total Harga : %d\n", totalHarga)


		totalPendapatan += totalHarga
	}

	fmt.Println("Total Pendapatan : ", totalPendapatan)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_lingkaran](/praktikum/05-perulangan-for-loop/guided/guided-3/output%20Bioskop.png)

#### Deskripsi
Pada program `bioskop.go`, saya membuat program untuk menghitung harga tiket bioskop berdasarkan jenis film, hari pemutaran, dan jumlah tiket. Program menentukan harga tiket, menghitung total pembayaran, dan memberikan diskon 10% jika membeli minimal lima tiket. Dengan perulangan `for`, program memproses setiap transaksi dan menghitung total pendapatan dari seluruh transaksi.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [prima.go]

```go
package main

import "fmt"

func main() {
	var n int
	var jumlah int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
			jumlah++
		}
	}

	fmt.Println()

	if jumlah == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_cacahuang](/praktikum/05-perulangan-for-loop/unguided/unguided-1/output%20prima.png)


#### Deskripsi
Pada program `prima.go`, saya membuat program untuk menampilkan semua faktor dari sebuah bilangan dan menentukan apakah bilangan tersebut termasuk bilangan prima. Program menggunakan `perulangan for` untuk memeriksa bilangan dari 1 sampai nilai yang dimasukkan. Jika bilangan habis dibagi, bilangan tersebut ditampilkan sebagai faktor. Setelah itu, program menentukan apakah jumlah faktornya tepat dua untuk mengetahui apakah bilangan tersebut prima.

### 2. [ecct.go]

```go

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
```

##### Output
![Output_kalkulator](/praktikum/05-perulangan-for-loop/unguided/unguided-2/output%20ecct.png)

#### Deskripsi
Pada program `ecct.go`, saya membuat program untuk menentukan apakah mahasiswa mendapatkan sertifikat ECCT berdasarkan rata-rata enam nilai tes. Program menghitung rata-rata nilai setiap mahasiswa. Jika rata-ratanya minimal 3,33, mahasiswa dinyatakan mendapat sertifikat. Setelah semua mahasiswa diperiksa menggunakan perulangan for, program menampilkan jumlah mahasiswa yang mendapatkan sertifikat dan tidak mendapatkan sertifikat.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->

## Kesimpulan
Setelah melakukan praktikum Modul 05, saya memahami penggunaan perulangan for dalam bahasa Go untuk menjalankan instruksi secara berulang. Melalui program yang dikerjakan, saya belajar menggunakan perulangan untuk menghitung total tabungan, mencari bilangan terbesar, menentukan faktor dan bilangan prima, serta menghitung jumlah mahasiswa yang mendapatkan sertifikat ECCT. Saya juga memahami bahwa perulangan dapat digunakan bersama variabel dan percabangan untuk memproses data sesuai kondisi yang ditentukan. Dengan praktikum ini, saya menjadi lebih memahami cara kerja perulangan for dan penerapannya dalam menyelesaikan permasalahan menggunakan bahasa Go.

## Referensi

1. Universitas Telkom. (t.t.). Modul 5: Perulangan For Loop. Materi Praktikum Algoritma dan Pemrograman, Laboratorium Praktikum Informatika, Fakultas Informatika, Program Studi S1 Rekayasa Perangkat Lunak, hlm. 28–31.
2. Universitas Telkom. (t.t). Soal Tugas Praktikum Modul 05: Algoritma dan Pemrograman — Perulangan For Bahasa Pemrograman Go (Golang). Laboratorium Praktikum Informatika, Fakultas Informatika, Program Studi S1 Rekayasa Perangkat Lunak, Semester Ganjil Tahun Akademik 2025/2026.
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->