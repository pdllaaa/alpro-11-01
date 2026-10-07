# <h1 align="center">Tugas Pendahuluan Modul 04 - “RUNTUNAN/SEKUENSI” Bahasa Pemrograman Go (Golang)</h1>
<p align="center">Fadila Riska Amalia Juniar - 109092600007</p>

### 1. ekspresi_control.go

```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10 
	var sngNum float64 = -3

	fmt.Println(intNum > 5)
	fmt.Println(intNum >= 5 && intOther < 11)
	fmt.Println(sngNum != -1 || intOther < 0)
	fmt.Println((!(intNum > 3) || intNum <= 5 ))
	fmt.Println((!(intOther >= intNum) ))
	fmt.Println((0 - sngNum > 0 ))
	fmt.Println((4 / 2 == intOther / intNum))
	fmt.Println((intOther % 2 == 0 ))
	fmt.Println((intOther + 2 * intNum != 30 || !(sngNum > 0) ))
	fmt.Println((intOther > 0 && intNum > 0 || sngNum > 0 ))
	fmt.Println((sngNum > 0 || (intNum >= 0 && -1 * intOther == -10) ))
	fmt.Println((intNum == 5))
	fmt.Println((intNum > 0 || (sngNum <= 0 && intOther == 13) ))
	fmt.Println((!(!(!(!(intNum > 0)))) ))
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/ekspresi/output%20ekspresi_control.png)


#### Deskripsi
Pada program `ekspresi_control`, saya menggunakan tiga variabel yaitu **intNum**, **intOther**, dan **sngNum**. Setelah itu, saya membuat beberapa kondisi menggunakan operator perbandingan dan operator logika seperti **>, >=, |=, &, ||, dan |**. Setiap kondisi akan menghasilkan output dengan nilai **true** atau **false**.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya belajar cara mengecek suatu kondisi menggunakan operator perbandingan dan logika di bahasa Go. Hasil dari setiap kondisi kemudian ditampilkan menggunakan `fmt.Println`.

### 2. tracing.go

```go
package main  

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}
	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}
	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}
	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}
	fmt.Println("Nilai akhir result:", result)
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/tracing/output%20tracing.png)


#### Deskripsi
Pada program tracing, saya menggunakan variabel x, y, z, dan result. Program ini memiliki beberapa kondisi if, else if, dan else yang dijalankan secara berurutan. Setiap kondisi akan memengaruhi nilai dari variabel result.

Pada kondisi pertama, x > 5 bernilai true dan y < 10 juga bernilai true. Jadi, result diisi dengan x + y, yaitu 15.
Selanjutnya, kondisi z > 10 && x == 10 bernilai true, sehingga result ditambah dengan z. Nilai result menjadi 30.

Pada kondisi ketiga, x == 10 || y > 10 bernilai true, sehingga result ditambah 5 dan menjadi 35.
Pada kondisi terakhir, x < 15 && y < 10 bernilai true. Karena terdapat tanda !, hasilnya menjadi false, sehingga bagian else dijalankan. Nilai result dikurangi 10 dan menjadi 25. Jadi, nilai akhir result adalah 25 dan output yang dihasilkan adalah:
Nilai akhir result: 25

Program ini merupakan bagian dari Tugas Pendahuluan. Dari program ini, saya belajar memahami bagaimana kondisi if, else if, dan else dijalankan serta bagaimana kondisi tersebut dapat mengubah nilai sebuah variabel.

1. Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi? 25

2. Apa output yang dihasilkan oleh program? Nilai akhir result: 25

3. Tuliskan langkah-langkah alur eksekusi program berdasarkan kondisi yang diberikan:

	◦ Kondisi 1: Apakah kondisi x > 5 benar? Jika ya, apa yang terjadi selanjutnya?
	
	Kondisi 1: x > 5 → true, kemudian y < 10 → true. Jadi result = 10 + 5 = 15.
	
	◦ Kondisi 2: Apakah kondisi z > 10 && x == 10 benar? Bagaimana hal ini memengaruhi nilai result?
	
	Kondisi 2: z > 10 && x == 10 → true. Jadi result = 15 + 15 = 30.

  	◦ Kondisi 3: Apakah salah satu dari kondisi x == 10 || y > 10 benar? Apa yang terjadi?
	
	Kondisi 3: x == 10 || y > 10 → true. Jadi result = 30 + 5 = 35.

  	◦ Kondisi 4: Bagaimana kondisi !(x < 15 && y < 10) dievaluasi? Apa dampaknya pada result?

	Kondisi 4: x < 15 && y < 10 → true, tetapi karena menggunakan !, hasilnya menjadi false. Jadi masuk else dan result = 35 - 10 = 25.

Program ini merupakan bagian dari **Tugas Pendahuluan**. Dari program ini, saya mempraktikkan cara menerima input, melakukan perhitungan sederhana, dan menampilkan hasil menggunakan `fmt.Println`

### 3. jumlahHari_sebulan.go

```go
package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun, &bulan)

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)

	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)

	case "Feb":
		if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		fmt.Println("Nama Bulan tidak valid")
   }
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/jumlah%20hari/output%20jumlahHari_sebulan.png)


#### Deskripsi
Pada program `jumlahHari_sebulan`, saya memasukkan tahun dan nama bulan menggunakan tiga huruf pertama, seperti `Jan`, `Feb`, dan `Mar`. Program menggunakan `switch case` untuk menentukan jumlah hari dari bulan yang dimasukkan. Untuk bulan Februari, program juga menggunakan kondisi untuk mengecek apakah tahun tersebut merupakan tahun kabisat atau bukan. Jika nama bulan tidak sesuai, program akan menampilkan pesan bahwa nama bulan tidak valid.

Program ini merupakan bagian dari Tugas Pendahuluan. Dari program ini, saya belajar menggunakan switch case, kondisi if, serta operator untuk mengecek tahun kabisat dalam bahasa Go

### 4. hari.go

```go
package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	switch angka {
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
		fmt.Println("Nomor hari tidak valid", angka)
	}
}

```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/swicth%20hari/output%20hari.png)


#### Deskripsi
Pada program `hari`, saya memasukkan angka dari 1 sampai 7 untuk menentukan nama hari. Program menggunakan `switch case` untuk mencocokkan angka dengan nama hari, mulai dari Senin sampai Minggu. Jika angka yang dimasukkan tidak sesuai dengan pilihan 1 sampai 7, program akan menampilkan pesan bahwa nomor hari tidak valid. Dari program ini, saya belajar menggunakan `switch case` untuk memilih hasil berdasarkan nilai input.

Program ini merupakan bagian dari Tugas Pendahuluan. Dari program ini, saya belajar menggunakan switch case untuk memilih hasil berdasarkan nilai yang dimasukkan.

## Kesimpulan
Setelah mengerjakan Tugas Pendahuluan Modul 04, saya lebih memahami penggunaan percabangan dan kondisi dalam bahasa Go. Saya belajar menggunakan operator perbandingan dan operator logika untuk menentukan nilai `true` atau `false`. Saya juga belajar memahami alur penggunaan `if`, `else if`, dan `else`, serta menggunakan `switch case` untuk menentukan hasil berdasarkan input. Melalui beberapa program tersebut, saya dapat memahami bagaimana kondisi dalam program dapat memengaruhi hasil yang ditampilkan.