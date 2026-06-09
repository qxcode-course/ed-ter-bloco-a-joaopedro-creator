package main
import "fmt"

func rainhas(linha, n int, colunas []int) int {
    if linha == n {
        return 1
    }
    
    count := 0
    for col := 0; col < n; col++ {
        valido := true
        for l := 0; l < linha; l++ {
            if colunas[l] == col || abs(colunas[l]-col) == abs(l-linha) {
                valido = false
                break
            }
        }
        if valido {
            colunas[linha] = col
            count += rainhas(linha+1, n, colunas)
        }
    }
    return count
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    var n int 
    fmt.Scan(&n)
    colunas := make([]int, n)
    fmt.Println(rainhas(0, n, colunas))
}