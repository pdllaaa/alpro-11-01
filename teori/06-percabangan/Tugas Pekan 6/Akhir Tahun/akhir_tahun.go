package main

import "fmt"

func main() {
    var belanja, diskon, cashback, bayar int
    var bersedia, kartu bool
    var dapatDiskon, dapatCashback bool

    fmt.Scan(&belanja, &bersedia)

    kartu = bersedia
    diskon = 0
    cashback = 0
    dapatDiskon = false
    dapatCashback = false

    if belanja >= 100000 {
        dapatDiskon = true
        diskon = belanja * 10 / 100
    }

    bayar = belanja - diskon

    if kartu && belanja >= 200000 {
        dapatCashback = true
        cashback = 75000
    }

    bayar = bayar - cashback

    fmt.Println("Kartu?", kartu)
    fmt.Println("Diskon?", dapatDiskon)
    fmt.Println("Cashback?", dapatCashback)
    fmt.Printf("Rp. %d\n", bayar)
}