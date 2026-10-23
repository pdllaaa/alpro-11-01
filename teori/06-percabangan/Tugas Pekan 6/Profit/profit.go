package main

import "fmt"

func main() {
    var sebelumnya, sekarang, selisih float64

    fmt.Scan(&sebelumnya, &sekarang)

    if sekarang > sebelumnya {
        selisih = sekarang - sebelumnya
        fmt.Printf("Peningkatan sebesar %g\n", selisih)
    } else if sekarang < sebelumnya {
        selisih = sebelumnya - sekarang
        fmt.Printf("Penurunan sebesar %g\n", selisih)
    } else {
        fmt.Println("Tetap")
    }
}