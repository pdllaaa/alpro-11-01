# <h1 align="center">Laporan Praktikum Modul 06 - Perulangan For Bahasa Pemrograman Go (Golang) 
</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

## Dasar Teori

### A. Pengertian Perulangan (Looping)
Perulangan atau looping adalah proses menjalankan perintah yang sama secara berulang selama kondisi tertentu masih terpenuhi. Perulangan membantu membuat program menjadi lebih sederhana karena kita tidak perlu menulis perintah yang sama berkali-kali. Contohnya adalah menghitung bilangan, memeriksa data, dan meminta pengguna memasukkan password sampai benar atau kesempatan habis.

### B. Perulangan While dalam Bahasa Go
Perulangan while digunakan untuk menjalankan perintah selama kondisi bernilai benar (true). Dalam bahasa Go, tidak terdapat kata kunci while secara langsung. Oleh karena itu, perulangan while ditulis menggunakan for yang hanya memiliki kondisi. Program akan terus menjalankan perintah di dalam perulangan selama kondisinya masih benar dan berhenti ketika kondisi bernilai salah (false).

#### C. Bentuk dan Cara Kerja Perulangan While
Bentuk perulangan while dalam Go adalah for kondisi { ... }. Sebelum menjalankan perintah di dalamnya, program akan memeriksa kondisi terlebih dahulu. Jika kondisi benar, perintah akan dijalankan dan kondisi diperiksa kembali. Proses ini terus berulang sampai kondisi tidak terpenuhi. Oleh karena itu, variabel yang digunakan dalam kondisi perlu diperbarui agar perulangan dapat berhenti pada waktu yang tepat.

#### D. Perulangan Repeat-Until
Perulangan repeat-until adalah perulangan yang menjalankan perintah terlebih dahulu, kemudian memeriksa kondisi untuk menentukan apakah perulangan perlu dihentikan. Dengan cara ini, perintah di dalam perulangan dijalankan setidaknya satu kali. Bahasa Go tidak memiliki kata kunci repeat-until secara langsung, sehingga logika tersebut perlu dibuat menggunakan struktur perulangan dan percabangan yang sesuai.

#### E. Penggunaan Perulangan dalam Program
Perulangan dapat digunakan untuk menyelesaikan berbagai masalah, seperti mencari bilangan yang memenuhi kondisi tertentu, menghitung jumlah huruf vokal dalam kalimat, dan membuat simulasi ATM. Pada program ATM, perulangan digunakan untuk mengulang proses login selama kesempatan masih tersedia dan menampilkan menu selama pengguna belum memilih keluar. Dengan menggabungkan perulangan, variabel, dan percabangan, program dapat menjalankan instruksi sesuai kondisi yang telah ditentukan.

## Guided

### 1. [terkecil.go]

```go
package main

import "fmt"

func main() {
	var batas int
	fmt.Scan(&batas)

	n := 1
	for n*n*n <= batas {
		n++
	}

	fmt.Println(n)
}	
```
#### Deskripsi
Pada program `terkecil.go`, saya membuat program untuk mencari bilangan bulat terkecil yang nilai pangkat tiganya lebih besar dari batas yang dimasukkan. Program dimulai dari nilai `n = 1`, kemudian menggunakan perulangan for untuk menaikkan nilai n selama nilai pangkat tiganya masih kurang dari atau sama dengan batas. Setelah kondisi tidak terpenuhi, program menampilkan nilai n yang memenuhi syarat.

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_terkecil](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-1/output%20terkecil.png)

### 2. [terbesar.go]

```go
package main

import "fmt"

func main() {
	var batas int
	fmt.Scan(&batas)

	n := 1
	for n*n < batas {
		n++
	}

	fmt.Println(n - 1)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_terbesar](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-2/output%20terbesar%201.png)

#### Deskripsi
Pada program `terbesar.go`, saya membuat program untuk mencari bilangan bulat terbesar yang nilai kuadratnya masih lebih kecil dari batas yang dimasukkan. Program dimulai dari `n = 1`, lalu menaikkan nilai n selama nilai kuadratnya masih lebih kecil dari batas. Setelah perulangan berhenti, program menampilkan nilai n - 1 sebagai bilangan terbesar yang memenuhi syarat.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. [vokal.go]

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    kalimat, _ := reader.ReadString('\n')

    i := 0
    jumlahVokal := 0

    for i < len(kalimat) {
        ch := strings.ToLower(string(kalimat[i]))
        if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
            jumlahVokal++
        }
        i++
    }

    fmt.Println(jumlahVokal)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_vokal](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-3/output%20vokal.png)

#### Deskripsi
Pada program `vokal.go`, saya membuat program untuk menghitung jumlah huruf vokal dalam sebuah kata atau kalimat yang dimasukkan oleh pengguna. Program memeriksa setiap karakter menggunakan perulangan for dan mengubah huruf menjadi huruf kecil agar huruf kapital juga dapat dihitung. Jika karakter merupakan huruf a, i, u, e, atau o, jumlah vokal akan bertambah satu. Setelah semua karakter diperiksa, program menampilkan jumlah huruf vokal yang ditemukan.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [atm.go]

```go

package main

import "fmt"

func main() {
	var password string
	var percobaan int
	var saldo int = 10000000
	var pilihan int
	var jumlah int
	var login bool = false

	// Proses login ATM
	for percobaan < 3 && !login {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == "1234" {
			login = true
			fmt.Println("Login berhasil!")
		} else {
			percobaan++
			fmt.Println("Password salah. Silakan coba lagi.")
		}
	}

	// Jika password salah sebanyak 3 kali
	if !login {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	// Menu ATM
	for pilihan != 3 {
		fmt.Println("\nMenu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)

		case 2:
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&jumlah)

			if jumlah <= 0 {
				fmt.Println("Jumlah penarikan harus lebih dari 0.")
			} else if jumlah <= saldo {
				saldo -= jumlah
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", jumlah)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}

		case 3:
			fmt.Println("Keluar dari sistem. Terima kasih!")

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Output_atm](/praktikum/06-perulangan-while-loop-dan-repeat/unguided/unguided-1/output%20atm.png)

#### Deskripsi
Pada program `atm.go`, saya membuat simulasi ATM yang memiliki fitur login, melihat saldo, dan menarik uang. Pengguna diberi maksimal tiga kesempatan untuk memasukkan password yang benar. Jika password salah sebanyak tiga kali, ATM akan terblokir. Jika login berhasil, menu ATM akan ditampilkan berulang kali menggunakan perulangan for sampai pengguna memilih keluar. Saat melakukan penarikan, program memeriksa apakah saldo mencukupi sebelum mengurangi saldo.

## Kesimpulan
Setelah melakukan praktikum Modul 06, saya memahami penggunaan perulangan while dalam bahasa Go dengan menggunakan for yang hanya memiliki kondisi. Saya belajar menggunakan perulangan untuk mencari bilangan yang memenuhi syarat, menghitung jumlah huruf vokal, dan membuat simulasi ATM. Saya juga memahami bahwa perulangan dapat digabungkan dengan variabel dan percabangan untuk menjalankan perintah sesuai kondisi tertentu. Melalui praktikum ini, saya menjadi lebih memahami cara kerja perulangan dan penerapannya dalam membuat program.

## Referensi

1. Universitas Telkom. (t.t.). Modul 6: Logika While & Repeat-Until. Materi Praktikum Algoritma dan Pemrograman, Laboratorium Praktikum Informatika, Fakultas Informatika, Program Studi S1 Rekayasa Perangkat Lunak.
2. Universitas Telkom. (t.t.). Soal Tugas Praktikum Modul 06: Algoritma dan Pemrograman — Perulangan While Bahasa Pemrograman Go (Golang). Laboratorium Praktikum Informatika, Fakultas Informatika, Program Studi S1 Rekayasa Perangkat Lunak.
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->