# <h1 align="center">Laporan Praktikum Modul 04 - “RUNTUNAN/SEKUENSI” Bahasa Pemrograman Go (Golang)
</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

## Dasar Teori

### A. Pengertian Runtunan atau Sekuensi
Runtunan atau sekuensi adalah cara menjalankan instruksi program secara berurutan, mulai dari baris pertama hingga terakhir. Setiap instruksi dijalankan satu kali sesuai urutan penulisannya. Urutan ini penting karena dapat memengaruhi hasil akhir program.

### B. Penggunaan Variabel dalam Program
Variabel digunakan untuk menyimpan data yang dibutuhkan dalam program. Dalam bahasa Go, tipe data int digunakan untuk bilangan bulat, float64 untuk bilangan desimal, dan string untuk teks. Variabel dapat menyimpan data yang dimasukkan pengguna maupun hasil perhitungan.

#### C. Input menggunakan fmt.Scan
Input adalah data yang dimasukkan pengguna ke dalam program. Dalam Go, fmt.Scan() digunakan untuk membaca masukan dan menyimpannya ke dalam variabel. Tanda & digunakan agar data yang dimasukkan dapat disimpan pada variabel yang dituju.

#### D. Output menggunakan fmt.Println dan fmt.Printf
Output adalah hasil yang ditampilkan setelah program selesai memproses data. Fungsi fmt.Println() digunakan untuk menampilkan teks atau nilai ke layar, sedangkan fmt.Printf() digunakan untuk menampilkan hasil dengan format tertentu, misalnya mengatur jumlah angka di belakang koma

#### E. Operator Aritmatika
Operator aritmatika digunakan untuk melakukan perhitungan dalam program. Operator yang umum digunakan adalah penjumlahan (+), pengurangan (-), perkalian (*), pembagian (/), dan sisa bagi (%). Operator tersebut dapat digunakan untuk menghitung nilai, pajak, maupun berbagai perhitungan lainnya.

#### F. Urutan Instruksi dan Proses Perhitungan
Dalam program, data biasanya dimasukkan terlebih dahulu, kemudian diolah, lalu hasilnya ditampilkan. Jika perhitungan terdiri dari beberapa tahap, hasil dari tahap sebelumnya dapat digunakan pada tahap berikutnya. Oleh karena itu, urutan instruksi harus diperhatikan agar hasil program sesuai dengan yang diharapkan.

## Guided

### 1. [grade.go]

```go
package main

import (
	"bufio"
	"os"

	"fmt"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
```
#### Deskripsi
Pada program grade.go, saya membuat program untuk menentukan grade berdasarkan nilai siswa. Program membaca nama dan nilai yang dimasukkan, kemudian memeriksa nilai tersebut menggunakan percabangan if dan else if. Setelah itu, program menentukan grade A, B, C, D, atau F sesuai rentang nilai dan menampilkan nama beserta grade yang diperoleh menggunakan `fmt.Printf`.

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_skor](/praktikum/04-runtunan-sekuensi/guided/soal-1/output%20grade.png)

### 2. [penilaian.go]

```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_tukar](/praktikum/04-runtunan-sekuensi/guided/soal-2/output%20penilaian.png)

#### Deskripsi
Pada program penilaian.go, saya membuat program penilaian sederhana menggunakan menu pilihan. Pengguna dapat memilih untuk menjalankan sistem penilaian atau keluar dari program. Jika memilih sistem penilaian, program meminta nama dan nilai siswa, kemudian menentukan grade berdasarkan nilai tersebut. Jika memilih keluar atau memasukkan pilihan yang tidak sesuai, program akan menampilkan pesan yang sesuai.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. [klasifikasi.go]

```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_lingkaran](/praktikum/04-runtunan-sekuensi/guided/soal-3/output%20klasifikasi.png)

#### Deskripsi
Pada program klasifikasi.go, saya membuat program untuk menentukan keterangan berdasarkan usia dan gaji yang dimasukkan pengguna. Program membaca kedua data tersebut, kemudian menggunakan switch untuk memeriksa kondisi yang sesuai. Hasilnya berupa keterangan seperti "Masih sekolah", "Muda sukses", atau "Pekerja mapan", tergantung pada usia dan gaji pengguna.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [grade.go]

```go
package main

import (
	"bufio"
	"os"

	"fmt"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80 && nilai < 90:
		grade = "B"
	case nilai >= 70 && nilai < 80:
		grade = "C"
	case nilai >= 60 && nilai < 70:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_cacahuang](/praktikum/04-runtunan-sekuensi/unguided/soal%201/output%20grade1.png)


#### Deskripsi
Pada program grade.go bagian Unguided, saya membuat program untuk menentukan grade siswa berdasarkan nilai yang dimasukkan. Program membaca nama dan nilai siswa, kemudian menggunakan switch untuk menentukan grade A, B, C, D, atau F sesuai rentang nilai. Setelah proses selesai, program menampilkan nama siswa beserta grade yang diperoleh.

### 2. [pajak.go]

```go
package main

import "fmt"

func main() {
	var penghasilan, pajak float64

	fmt.Scan(&penghasilan)

	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
	} else if penghasilan <= 100 {
		pajak = 0.05*50 + 0.10*(penghasilan-50)
	} else if penghasilan <= 200 {
		pajak = 0.05*50 + 0.10*50 + 0.15*(penghasilan-100)
	} else {
		pajak = 0.05*50 + 0.10*50 + 0.15*100 + 0.20*(penghasilan-200)
	}

	fmt.Println(pajak)
}
```

##### Output
![Output_kalkulator](/praktikum/04-runtunan-sekuensi/unguided/soal%202/output%20pajak.png)

#### Deskripsi
Pada program pajak.go, saya membuat program untuk menghitung pajak berdasarkan jumlah penghasilan yang dimasukkan pengguna. Program membaca penghasilan, kemudian menghitung pajak menggunakan beberapa kondisi sesuai rentang penghasilan. Setiap rentang memiliki tarif pajak yang berbeda, yaitu 5%, 10%, 15%, dan 20%. Setelah perhitungan selesai, program menampilkan jumlah pajak yang harus dibayar menggunakan `fmt.Println`.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->

## Kesimpulan
Setelah melakukan praktikum ini, saya dapat memahami cara kerja program menggunakan bahasa Go, mulai dari penggunaan variabel, tipe data, input dan output, hingga percabangan untuk mengolah data. Saya juga belajar membuat program untuk menentukan grade siswa, mengklasifikasikan data berdasarkan usia dan gaji, serta menghitung pajak berdasarkan penghasilan. Melalui praktikum ini, saya menjadi lebih memahami bagaimana menyusun kode dan menggunakan kondisi yang sesuai agar program dapat berjalan dan menghasilkan output yang diharapkan.

## Referensi

1. Universitas Telkom. (t.t.). Modul 4: Runtunan/Sekuensi. Materi Praktikum Algoritma dan Pemrograman, hlm. 17–27.
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->