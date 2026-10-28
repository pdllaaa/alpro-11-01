# <h1 align="center">Tugas Pendahuluan Modul 07 - “PERCABANGAN” Bahasa Pemrograman Go (Golang)</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

### 1. Jelaskan perbedaan antara if, if-else, else if, dan switch pada bahasa pemrograman Go! Sebutkan juga dua perbedaan switch pada Go dibandingkan dengan switch pada Java. 

#### Jawaban
- `if` digunakan untuk menjalankan kode jika suatu kondisi bernilai benar.
- `if-else` digunakan untuk memilih satu dari dua kemungkinan berdasarkan kondisi.
- `else if` digunakan untuk memeriksa beberapa kondisi secara berurutan.
- `switch` digunakan untuk memilih kode berdasarkan nilai atau kondisi yang sesuai dengan case.

Dua perbedaan switch pada Go dibandingkan dengan Java:
1. Pada Go, switch secara otomatis berhenti setelah satu case cocok, sehingga tidak memerlukan break. Pada Java, switch tradisional umumnya memerlukan break untuk mencegah fall-through.
2. Pada Go, switch dapat digunakan tanpa ekspresi, sehingga setiap case bisa berisi kondisi Boolean. Pada Java, switch tradisional menggunakan ekspresi yang nilainya dicocokkan dengan case. 

### 2. bilangan.go

```go
package main

import "fmt"

func main() {
    var bilangan int
    fmt.Scan(&bilangan)

    if bilangan == 0 {
        fmt.Println("Nol")
        fmt.Println("0 adalah bilangan nol")
    } else if bilangan%2 == 0 {
        if bilangan > 0 {
            fmt.Println("Genap, Positif")
        } else {
            fmt.Println("Genap, Negatif")
        }
    } else {
        if bilangan > 0 {
            fmt.Println("Ganjil, Positif")
        } else {
            fmt.Println("Ganjil, Negatif")
        }
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output bilangan](/praktikum/07-percabangan/tp/bilangan/output%20Bilangan.png)


#### Deskripsi
Pada program `bilangan.go`, saya menggunakan variabel bilangan untuk menyimpan bilangan bulat yang dimasukkan oleh user. Selanjutnya, program memeriksa apakah bilangan bernilai 0, genap, atau ganjil, lalu menentukan apakah bilangan tersebut positif atau negatif. Setelah itu, program menampilkan hasil sesuai dengan kondisi bilangan yang dimasukkan.

Program ini merupakan bagian dari `Tugas Pendahuluan`. Dari program ini, saya belajar cara menggunakan percabangan `if, else if, dan else` dalam bahasa Go untuk memeriksa kondisi dan menentukan keluaran program.

### 3. hari.go

```go
package main

import "fmt"

func main() {
    var hari int
    fmt.Scan(&hari)

    switch hari {
    case 1:
        fmt.Println("Senin")
    case 2:
        fmt.Println("Selasa")
    case 3:
        fmt.Println("Rabu")
    case 4:
        fmt.Println("Kamis")
    case 5:
        fmt.Println("Jumat")
    case 6:
        fmt.Println("Sabtu")
    case 7:
        fmt.Println("Minggu")
    default:
        fmt.Println("Nomor hari tidak valid")
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output hari](/praktikum/07-percabangan/tp/hari/output_Hari.png)

#### Deskripsi
Pada program `hari.go`, saya menggunakan variabel hari untuk menyimpan nomor hari yang dimasukkan oleh user. Selanjutnya, program menggunakan `switch` untuk menentukan nama hari sesuai nomor yang dimasukkan. Jika nomor hari adalah 6 atau 7, program menambahkan keterangan akhir pekan. Jika nomor di luar 1–7, program menampilkan pesan bahwa nomor hari tidak valid.

Program ini merupakan bagian dari `Tugas Pendahuluan`. Dari program ini, saya belajar menggunakan percabangan `switch` untuk menentukan keluaran berdasarkan nilai yang dimasukkan.

### 4. terbesar.go

```go
package main

import "fmt"

func main() {
    var a, b, c int
    fmt.Scan(&a, &b, &c)

    if a >= b && a >= c {
        fmt.Println(a)
    } else if b >= a && b >= c {
        fmt.Println(b)
    } else {
        fmt.Println(c)
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output terbesar](/praktikum/07-percabangan/tp/terbesar/output_Terbesar.png)


#### Deskripsi
Pada program `terbesar.go`, saya menggunakan variabel `a, b, dan c` untuk menyimpan tiga bilangan bulat yang dimasukkan oleh user. Selanjutnya, program membandingkan ketiga bilangan menggunakan `if-else` untuk menentukan bilangan yang paling besar. Setelah itu, program menampilkan bilangan terbesar sebagai hasil akhir.

Program ini merupakan bagian dari `Tugas Pendahuluan`. Dari program ini, saya belajar menggunakan percabangan `if-else` dan `operator perbandingan` untuk menentukan nilai terbesar dari tiga bilangan.

## Kesimpulan
Setelah mengerjakan Tugas Pendahuluan Modul 07, saya menjadi lebih memahami penggunaan percabangan dalam bahasa Go. Saya belajar menggunakan if, else if, else, dan switch untuk memeriksa kondisi, menentukan nama hari, serta mencari bilangan terbesar. Dari tugas ini, saya jadi lebih paham cara menggunakan percabangan untuk menyelesaikan masalah dalam program.