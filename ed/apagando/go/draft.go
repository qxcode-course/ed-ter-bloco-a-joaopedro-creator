package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    fila := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&fila[i])
    }

    var m int
    fmt.Scan(&m)

    saiu := make([]int, m)
    for i := 0; i < m; i++ {
        fmt.Scan(&saiu[i])
    }

    removidos := make(map[int]bool)

    for i := 0; i < m; i++ {
        removidos[saiu[i]] = true
    }

    for i := 0; i < n; i++ {
        if !removidos[fila[i]] {
            fmt.Printf("%d ", fila[i])
        }
    }

    fmt.Println()
}