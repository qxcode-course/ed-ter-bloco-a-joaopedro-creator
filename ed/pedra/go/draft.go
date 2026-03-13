package main
import "fmt"
import "math"
func main() {
    var N int 
    fmt.Scan(&N)

    menorPontuacao := 101
    vencedor := -1

    for i := 0; i < N ; i++ {
        var A, B int 
        fmt.Scan(&A, &B)

        if A < 10 || B < 10{
            continue
        }

        pontuacao := int(math.Abs(float64(A - B)))

        if pontuacao < menorPontuacao {
            menorPontuacao = pontuacao
            vencedor = i
        }
    }

    if vencedor == -1{
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(vencedor)
    }


}
