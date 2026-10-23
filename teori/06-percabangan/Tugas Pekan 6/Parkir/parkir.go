package main

import "fmt"

func main() {
    var h1, m1, h2, m2 int
    var mulai, selesai, durasi int
    var hh, mm int

    fmt.Scan(&h1, &m1, &h2, &m2)

    if h1 == 12 {
        h1 = 0
    }
    if h2 == 12 {
        h2 = 0
    }

    mulai = h1*60 + m1
    selesai = h2*60 + m2

    if selesai < mulai {
        selesai = selesai + 12*60
    }

    durasi = selesai - mulai
    hh = durasi / 60
    mm = durasi % 60

    fmt.Printf("%d jam %d menit\n", hh, mm)
}