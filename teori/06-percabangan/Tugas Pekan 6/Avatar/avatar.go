package main

import "fmt"

func main() {
    var p, dewasa, kecil int
    var sisa, kapasitas, tidakBerangkat int

    fmt.Scan(&p)

    dewasa = (p + 4) / 5

    if dewasa > 3 {
        dewasa = 3
    }

    kapasitas = dewasa * 5
    sisa = p - kapasitas
    kecil = 0

    if sisa > 0 {
        kecil = (sisa + 1) / 2

        if kecil > 5 {
            kecil = 5
        }
    }

    kapasitas = dewasa*5 + kecil*2
    tidakBerangkat = p - kapasitas

    fmt.Printf("Dewasa %d", dewasa)

    if kecil > 0 {
        fmt.Printf(", kecil %d", kecil)
    }

    if tidakBerangkat > 0 {
        fmt.Printf(", dan %d tak berangkat", tidakBerangkat)
    }

    fmt.Println()
}