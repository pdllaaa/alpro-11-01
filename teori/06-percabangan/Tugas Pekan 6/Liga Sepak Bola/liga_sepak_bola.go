package main

import "fmt"

func main() {
    var a, b, c, d int
    var terbesar, terkecil int

    fmt.Scan(&a, &b, &c, &d)

    terbesar = a
    terkecil = a

    if b > terbesar {
        terbesar = b
    }
    if c > terbesar {
        terbesar = c
    }
    if d > terbesar {
        terbesar = d
    }

    if b < terkecil {
        terkecil = b
    }
    if c < terkecil {
        terkecil = c
    }
    if d < terkecil {
        terkecil = d
    }

    fmt.Println(terbesar, terkecil)
}