# <h1 align="center">Tugas Pendahuluan Modul 05 - “PERULANGAN FOR” Bahasa Pemrograman Go (Golang)</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

### 1. Apa fungsi utama dari struktur perulangan for dalam program Go? Sebutkan dan jelaskan tiga komponen pada perulangan for (inisialisasi, kondisi, dan post statement)! 
Fungsi utama dari struktur perulangan `for` dalam pemrograman go adalah untuk menjalankan suatu blok kode secara berulang-ulang selama kondisi yang ditentukan masih bernilai true. Pada bentuk perulangan `for` yang memiliki tiga komponen, strukturnya sebagai berikut :

1. **Inisialisasi** merupakan bagian awal untuk menentukan nilai variabel yang digunakan pada perulangan. Bagian ini dijalankan hanya satu kali sebelum perulangan dimulai. Contohnya, `i := 1` berarti i dimulai dari angka 1.
2. **Kondisi** adalah bagian untuk menentukan apakah perulangan masih perlu dijalankan atau sudah harus berhenti. Selama kondisi bernilai benar `true` maka perulangan masih tetap berlanjut. Jika, kondisi bernilai salah `false` maka perulangan akan berhenti. Contohnya, `i <= 10` berarti perulangan berjalan selama nilai i masih kurang dari atau sama dengan 10.
3. **Post Statement** Post statement adalah bagian yang dijalankan setelah kode di dalam perulangan selesai dilakukan. Bagian ini biasanya digunakan untuk mengubah nilai variabel agar perulangan bisa mencapai batas akhirnya. Contohnya, i++ berarti nilai i bertambah satu setiap kali perulangan selesai dijalankan. 

### 2. Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh kode masing-masing.
1. **For dengan tiga komponen** merupakan perulangan yang memiliki nilai awal, kondisi, dan perubahan nilai variabel. Biasanya digunakan ketika jumlah perulangan sudah digunakan.
Contoh kode sebagai berikut :

``` go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}
```
Penjelasan dari kode tersebut adalah variabel `i` dimulai dari 1, perulangan akan berjalan selama `i` tidak lebih dari `5`, dan `i++` digunakan untuk menambah nilainya satu setiap perulangan. 

2. **For yang hanya menggunakan kondisi** merupakan sebuah perulangan yang hanya menuliskan kondisi pada `for`, sedangkan nilai awal dan perubahan variabel ditulis secara terpisah. Contoh kodenya sebagai berikut :

``` go
package main

import "fmt"

func main() {
    i := 1

    for i <= 5 {
        fmt.Println(i)
        i++
    }
}
```
Penjelasan dari kode tersebut adalah variabel `i` dimulai dari 1. Perulangan berjalan selama `i` tidak lebih dari `5`, kemudian `i++` menambah nilainya satu setiap perulangan.

**Kesimpulannya** terletak pada perbedaan dalam penulisan kode. `for` dengan `tiga komponen` menuliskan **nilai awal, kondisi, dan perubahan nilai dalam satu baris**, sedangkan `for` dengan `kondisi` saja menuliskan **kondisi pada bagian for dan perubahan nilainya di dalam blok perulangan**. Keduanya dapat menghasilkan output yang sama.

### 3. ganjil.go

```go
package main

import "fmt"

func main() {
    var n int
    var jumlah int = 0

    fmt.Scan(&n)

    for i := 1; i <= n; i += 2 {
        fmt.Print(i, " ")
        jumlah += i
    }

    fmt.Println()
    fmt.Println(jumlah)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/ganjil/output%20ganjil.png)


#### Deskripsi
Pada program `ganjil.go`, saya menggunakan variabel `int` untuk menyimpan bilangan bulat yang dimasukkan oleh user dan hasil penjumlahan bilangan ganjil. Setelah itu, saya menggunakan perulangan `for` untuk menampilkan bilangan ganjil dari `1 sampai batas angka yang dimasukkan`. Setiap bilangan ganjil akan dijumlahkan menggunakan variabel jumlah, kemudian hasil akhirnya ditampilkan menggunakan `fmt.Println`.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya belajar cara menggunakan perulangan for untuk menampilkan bilangan ganjil dan menghitung jumlah seluruh bilangan ganjil menggunakan bahasa Go.

### 4. perkalian.go

```go
package main

import "fmt"

func main() {
    var a, b int

    fmt.Scan(&a, &b)

    for i := 1; i <= b; i++ {
        fmt.Println(a, "x", i, "=", a*i)
    }
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/perkalian/output%20perkalian.png)


#### Deskripsi
Pada program `perkalian.go`, saya menggunakan variabel `int` untuk menyimpan bilangan yang akan dikalikan dan batas perkalian yang dimasukkan oleh user. Setelah itu, saya menggunakan perulangan `for` untuk menghitung perkalian mulai dari `1 sampai batas yang ditentukan`. Setiap hasil perkalian ditampilkan menggunakan `fmt.Println` dengan format yang mudah dibaca.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya belajar cara menggunakan perulangan for untuk melakukan perhitungan secara berulang dan menampilkan tabel perkalian menggunakan bahasa Go.

## Kesimpulan
Setelah mengerjakan Tugas Pendahuluan Modul 05, saya lebih memahami penggunaan perulangan `for` dalam bahasa Go. Saya belajar menggunakan perulangan dengan tiga komponen maupun kondisi saja, serta menerapkannya untuk menampilkan bilangan ganjil, menghitung jumlahnya, dan membuat tabel perkalian. Melalui tugas ini, saya menjadi lebih paham cara menggunakan perulangan untuk menjalankan perintah secara berulang dengan lebih sederhana dan efisien.