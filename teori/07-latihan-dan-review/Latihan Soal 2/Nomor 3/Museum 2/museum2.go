package main

import "fmt"

func main() {
    var pengunjung, total int

    total = 0

    for hari := 1; hari <= 5; hari++ {
        for {
            fmt.Scan(&pengunjung)

            if pengunjung >= 0 && pengunjung <= 200 {
                break
            }
        }

        total += pengunjung
    }

    fmt.Println("Jumlah pengunjung:", total)
}