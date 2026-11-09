package main

import "fmt"

func main() {
    var angka string
    var ascending, descending bool

    fmt.Scan(&angka)

    ascending = true
    descending = true

    for i := 0; i < len(angka)-1; i++ {
        if angka[i] >= angka[i+1] {
            ascending = false
        }

        if angka[i] <= angka[i+1] {
            descending = false
        }
    }

    if ascending {
        fmt.Println("ascending")
    } else if descending {
        fmt.Println("descending")
    } else {
        fmt.Println("tidak terurut")
    }
}