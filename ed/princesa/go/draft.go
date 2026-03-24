package main

import "fmt"

func main() {
    var n, e int
    fmt.Scan(&n, &e)

    vivos := make([]bool, n)

    // todos começam vivos
    for i := 0; i < n; i++ {
        vivos[i] = true
    }

    espada := e - 1
    qtdVivos := n

    for qtdVivos > 1 {

        // imprimir estado atual
        fmt.Print("[ ")
        for i := 0; i < n; i++ {
            if vivos[i] {
                if i == espada {
                    fmt.Printf("%d> ", i+1)
                } else {
                    fmt.Printf("%d ", i+1)
                }
            }
        }
        fmt.Println("]")

        // encontrar quem vai morrer
        morto := (espada + 1) % n
        for !vivos[morto] {
            morto = (morto + 1) % n
        }

        // matar
        vivos[morto] = false
        qtdVivos--

        // passar a espada
        espada = (morto + 1) % n
        for !vivos[espada] {
            espada = (espada + 1) % n
        }
    }

    // imprimir último sobrevivente
    fmt.Print("[ ")
    for i := 0; i < n; i++ {
        if vivos[i] {
            fmt.Printf("%d> ", i+1)
        }
    }
    fmt.Println("]")
}