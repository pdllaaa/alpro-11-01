
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