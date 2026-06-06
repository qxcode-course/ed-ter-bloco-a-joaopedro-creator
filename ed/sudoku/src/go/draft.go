package main
import (
    "fmt"
    "math"

)

func contem(arr []rune, val rune) bool {
    for _, v := range arr {
        if v == val {
            return true
        }
    }
    return false
}

func linha(matriz [][]rune, lin int) []rune {
    return matriz[lin]
}

func coluna(matriz [][]rune, col int) []rune {
    var c []rune
    for i := 0; i < len(matriz); i++ {
        c = append(c, matriz[i][col])
    }
    return c
}   

func quadrante(matriz [][]rune, lin, col int) []rune {
  dim := len(matriz)
tam := int(math.Sqrt(float64(dim)))
l := (lin / tam) * tam
c := (col / tam) * tam

    if dim == 4 {
        return []rune{
            matriz[l+0][c], matriz[l+0][c+1],
            matriz[l+1][c], matriz[l+1][c+1],
        }
    }

    if dim == 9 {
        return []rune{
            matriz[l+0][c], matriz[l+0][c+1], matriz[l+0][c+2],
            matriz[l+1][c], matriz[l+1][c+1], matriz[l+1][c+2],
            matriz[l+2][c], matriz[l+2][c+1], matriz[l+2][c+2],
        }
    }
    return nil
}



func resolver(matriz [][]rune, index int) bool {
    nl := len(matriz)
    l := index / nl
    c := index % nl
    if index == nl * nl {
        return true
    }
   
    // se não for ponto, continue
     if matriz[l][c] != '.' {
        return resolver(matriz, index+1)
    }
    // para todos os números de [1 a N]
    //     se o número não estiver na linha, coluna e quadrante
    //         coloque o número na matriz
    //         se resolver(matriz, index + 1):
    //             return true
    //         matriz[l][c] = '.' // desfaz a tentativa


    for num := '1'; num <= rune(nl)+'0'; num++{
        if !contem(linha(matriz, l), num) && !contem(coluna(matriz, c), num) && !contem(quadrante(matriz, l, c), num) {
            matriz[l][c] = num
            if resolver(matriz, index + 1) {
                return true
            }
            matriz[l][c] = '.'
        }
    }
     return false
}




func main() {
    var n int 
    fmt.Scan(&n)

    matriz := make([][]rune, n)
    for i := 0; i < n; i++ {
        var linha string
        fmt.Scan(&linha)
        matriz[i] = []rune(linha)
    }

    resolver(matriz, 0)
    for _, linha := range matriz {
        fmt.Println(string(linha))
    }
}