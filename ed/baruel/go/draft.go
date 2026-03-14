package main
import "fmt"

func main() {
    var total, n int
    fmt.Scan(&total, &n)

    fig := make([]int, n)
    tem := make([]bool, total+1)

    for i := 0; i < n; i++ {
        fmt.Scan(&fig[i])
        if fig[i] <= total {
            tem[fig[i]] = true
        }
    }

    repetidas := []int{}
    for i := 1; i < n; i++ {
        if fig[i] == fig[i-1] {
            repetidas = append(repetidas, fig[i])
        }
    }

    if len(repetidas) == 0 {
        fmt.Println("N")
    } else {
        fmt.Print(repetidas[0])
        for i := 1; i < len(repetidas); i++ {
            fmt.Print(" ", repetidas[i])
        }
        fmt.Println()
    }

    faltando := []int{}
    for i := 1; i <= total; i++ {
        if !tem[i] {
            faltando = append(faltando, i)
        }
    }

    if len(faltando) == 0 {
        fmt.Println("N")
    } else {
        fmt.Print(faltando[0])
        for i := 1; i < len(faltando); i++ {
            fmt.Print(" ", faltando[i])
        }
        fmt.Println()
    }
}