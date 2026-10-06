# <h1 align="center">Laporan Praktikum Modul 03 - Algoritma dan Pemrograman Variabel dan Operator Bahasa Pemrograman Go (Golang)
</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

## Dasar Teori

### A. Struktur Bahasa Pemrograman Go
Dalam bahasa pemrograman Go, program utama memiliki dua bagian penting, yaitu **package main dan func main()**. **Package main** menunjukkan bahwa file tersebut berisi program utama, sedangkan di **func main()** digunakan untuk menuliskan bagian utama dari program. Pada bahasa pemrograman Go juga menyediakan komentar untuk memberikan keterangan pada kode. **Komentar** dapat dibuat jadi satu baris menggunakan tanda **//**, sedangkan komentar beberapa baris menggunakan tanda /* dan */.

### B. Package dan Struktur Program di Go

#### 1. Pengertian Package main dan func main()
Pada program Go **Package main** digunakan sebagai penanda bahwa file tersebut merupakan bagian dari program utama. Sementara itu, **func main()** menjadi tempat untuk menuliskan perintah yang akan dijalankan ketika program dimulai. Pada contoh di Modul 2, package fmt digunakan untuk membantu proses input dan output.

#### 2. Koding, Kompilasi, dan Eksekusi Go
Program Go ditulis menggunakan penyunting teks dan disimpan dengan ekstensi **.go**. Satu program lengkap dapat terdiri dari beberapa file **.go** selama file-file tersebut berada dalam folder yang sama. Setelah program selesai dibuat, program dapat di kompilasi menggunakan perintah **go build** atau **go build file.go**. Jika proses kompilasi berhasil, akan terbentuk program yang dapat dijalankan. Go juga memiliki beberapa perintah lain seperti **go fmt** untuk merapikan format kode dan **go clean** untuk membersihkan file hasil kompilasi.

#### 3. Variabel dan Tipe Data
Variabel digunakan untuk menyimpan data yang akan digunakan dalam program. Setiap variabel memiliki tipe data sesuai dengan jenis data yang disimpan. Dalam praktikum ini digunakan beberapa tipe data, seperti int untuk menyimpan bilangan bulat dan float64 untuk menyimpan bilangan yang memiliki angka desimal. Contohnya, pada kode konversi suhu digunakan float64 karena suhu dapat memiliki nilai desimal, sedangkan pada program pecahan uang digunakan int karena jumlah uang yang dihitung berupa bilangan bulat.

#### 4. Input dan Output
Input digunakan untuk memasukkan data ke dalam program, sedangkan output digunakan untuk menampilkan hasil dari proses yang dilakukan. Pada bahasa pemrograman Go, input dapat dilakukan menggunakan fmt.Scan() dan fmt.Scanln(). Sementara itu, fmt.Println() digunakan untuk menampilkan hasil ke layar. Contohnya, pengguna memasukkan jumlah hari, kemudian program memprosesnya dan menampilkan hasil berupa tahun, bulan, minggu, dan hari.

#### 5. Operator Aritmatika
Operator aritmatika digunakan untuk melakukan perhitungan dalam program. Operator yang digunakan dalam praktikum ini adalah penjumlahan (+), pengurangan (-), perkalian (*), pembagian (/), dan sisa bagi (%). Operator tersebut digunakan sesuai dengan kebutuhan program, seperti menghitung konversi suhu, pecahan uang, dan jumlah hari.

#### 6. Pembagian dan Sisa Bagi
Pembagian (/) digunakan untuk mendapatkan hasil pembagian dua bilangan, sedangkan sisa bagi (%) digunakan untuk mendapatkan sisa dari pembagian. Kedua operator ini dapat digunakan secara bersamaan untuk memecah suatu nilai menjadi beberapa bagian. Contohnya pada program konversi jumlah hari, jumlah hari dibagi dengan 360 untuk mendapatkan jumlah tahun. Setelah itu, sisa pembagiannya digunakan untuk menghitung bulan, minggu, dan hari.

## Guided

### 1. [konversi.go]

```go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Masukkan suhu: ")
	fmt.Scanln(&celcius)

	fmt.Println(celcius + 273)
}
```
#### Deskripsi
Pada program konversi suhu, saya memasukkan suhu dalam satuan Celcius. Setelah itu, suhu tersebut ditambahkan dengan 273 untuk mendapatkan hasil dalam satuan Kelvin. Hasil yang ditampilkan berupa suhu dalam Kelvin.

Program ini merupakan program sederhana untuk mempraktikkan cara menerima input, melakukan perhitungan sederhana menggunakan operasi penjumlahan, dan menampilkan hasil menggunakan `fmt.Println`.

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_skor](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/konversi/outputkonversikelvin.png)

### 2. [tukar.go]

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_tukar](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/tukar/outputnilaitukar.png)

#### Deskripsi
Pada program pertukaran nilai, saya memasukkan tiga nilai berupa x, y, dan z. Setelah itu, nilai tersebut ditukar posisinya menggunakan variabel sementara temp, sehingga nilai x menjadi nilai z, nilai z menjadi nilai y, dan nilai y menjadi nilai x. Hasil yang ditampilkan berupa nilai x, y, dan z setelah ditukar.

Program ini digunakan untuk mempraktikkan cara menerima beberapa input, menggunakan variabel sementara untuk menukar nilai, dan menampilkan hasil menggunakan fmt.Println.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. [kasir.go]

```go
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var sepuluhRibu int = x / 10000
	var sisa int = x % 10000

	var limaRibu int = sisa / 5000
	sisa = sisa % 5000

	var seribu int = sisa / 1000

	fmt.Println(sepuluhRibu, limaRibu, seribu)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_lingkaran](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/kasir/outputkasir.png)

#### Deskripsi
Pada program pecahan uang, saya memasukkan jumlah uang yang ingin dihitung. Setelah itu, program menghitung jumlah pecahan uang Rp10.000, Rp5.000, dan Rp1.000 yang dapat diperoleh dari jumlah tersebut. Program menggunakan operasi pembagian dan sisa bagi untuk menentukan jumlah masing-masing pecahan. Hasil yang ditampilkan berupa jumlah lembar uang Rp10.000, Rp5.000, dan Rp1.000.

Program ini digunakan untuk mempraktikkan cara menerima input, menggunakan operasi pembagian dan sisa bagi, menyimpan hasil perhitungan dalam variabel, dan menampilkan hasil menggunakan `fmt.Println`.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [konversi_reamur.go]

```go
package main

import "fmt"

func main(){
	var celcius float64

	fmt.Scanln(&celcius)

	fmt.Println(celcius * 4 / 5)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_cacahuang](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi%20suhu/outputkonversireamur.png)


#### Deskripsi
Pada program konversi suhu, saya memasukkan suhu dalam satuan Celcius. Setelah itu, suhu tersebut dikonversi ke satuan Reamur dengan cara mengalikan nilai Celcius dengan 4 lalu membaginya dengan 5. Hasil yang ditampilkan berupa suhu dalam satuan Reamur.

Program ini digunakan untuk mempraktikkan cara menerima input, melakukan perhitungan menggunakan operasi perkalian dan pembagian, serta menampilkan hasil menggunakan `fmt.Println`.

### 2. [konversi_jumlahHari.go]

```go
package main

import "fmt"

func main() {
	var totalHari int

	fmt.Scan(&totalHari)

	// 1 tahun = 12 bulan x 30 hari = 360 hari
	tahun := totalHari / 360
	sisaHari := totalHari % 360

	//1 bulan = 30 hari
	bulan := sisaHari / 30
	sisaHari = sisaHari % 30

	//1 minggu = 7 hari
	minggu := sisaHari / 7
	sisaHari = sisaHari % 7

	//Output empat baris bilangan bulat
	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(minggu)
	fmt.Println(sisaHari)
}
```

##### Output
![Output_kalkulator](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi%20hari/outputkonversihari.png)

#### Deskripsi
Pada program konversi waktu, saya memasukkan jumlah hari. Setelah itu, program mengubah total hari tersebut menjadi tahun, bulan, minggu, dan sisa hari. Perhitungannya menggunakan ketentuan 1 tahun = 360 hari, 1 bulan = 30 hari, dan 1 minggu = 7 hari. Hasil yang ditampilkan berupa jumlah tahun, bulan, minggu, dan sisa hari secara berurutan.

Program ini digunakan untuk mempraktikkan cara menerima input, menggunakan operasi pembagian dan sisa bagi untuk melakukan konversi waktu, serta menampilkan hasil menggunakan `fmt.Println`.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->

## Kesimpulan
Setelah melakukan praktikum Modul 03, saya dapat memahami penggunaan variabel, tipe data, input dan output, serta operator aritmatika pada bahasa pemrograman Go. Saya juga dapat menerapkan pembagian dan sisa bagi dalam beberapa program, seperti konversi suhu, pertukaran nilai, pecahan uang, dan konversi jumlah hari.

## Referensi

1. The Go Authors. (n.d.). The Go Programming Language Specification: Comments. Diakses pada 6 Oktober 2026 melalui [https://go.dev/ref/spec#Comments].
2. The Go Authors. (n.d.). The Go Programming Language Specification: Operators and punctuation. Diakses pada 6 Oktober 2026 melalui [https://go.dev/ref/spec#Operators_and_punctuation]
3. The Go Authors. (n.d.). The Go Programming Language Specification: Numeric types. Diakses pada 6 Oktober 2026 melalui [https://go.dev/ref/spec#Numeric_types]
4. The Go Authors. (n.d.). How to Write Go Code. The Go Programming Language. Diakses pada 6 Oktober 2026 melalui [https://go.dev/doc/code]
5. TThe Go Authors. (n.d.). The Go Programming Language Specification: Variable declarations. Diakses pada 6 Oktober 2026 melalui [https://go.dev/ref/spec#Variable_declarations]
6. Universitas Telkom. (n.d.). Soal Tugas Praktikum Modul 03: Algoritma dan Pemrograman Variabel dan Operator Bahasa Pemrograman Go (Golang). Laboratorium Praktikum Informatika, Fakultas Informatika, Program Studi S1 Rekayasa Perangkat Lunak.
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->