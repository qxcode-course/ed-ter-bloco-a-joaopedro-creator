package main

import (
    "bufio"
    "fmt"
    "os"
)

var a []int
var n, k int

func backtrack(indice int, soma int) bool {
    if soma == k{
        return true
    }
    if indice == n{
        return false
    }
    if soma > k {
        return false
    }

    return backtrack(indice+1, soma) || backtrack(indice+1, soma + a[indice])

}



func main() {
   entrada := bufio.NewReader(os.Stdin)
   
   fmt.Fscan(entrada, &n, &k)

    a = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(entrada, &a[i])
    }

    if backtrack(0, 0) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
