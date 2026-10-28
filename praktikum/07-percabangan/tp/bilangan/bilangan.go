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