package main
import "fmt"

func main() {
    var Q int
    var D string
    fmt.Scan(&Q, &D)

    x := make([]int, Q)
    y := make([]int, Q)

    for i := 0; i < Q; i++ {
        fmt.Scan(&x[i], &y[i])
    }

    oldx := x[0]
    oldy := y[0]

    switch D {
    case "L":
        x[0]--
    case "R":
        x[0]++
    case "U":
        y[0]--
    case "D":
        y[0]++
    }

    for i := 1; i < Q; i++ {
        tmpx := x[i]
        tmpy := y[i]

        x[i] = oldx
        y[i] = oldy

        oldx = tmpx
        oldy = tmpy
    }

    for i := 0; i < Q; i++ {
        fmt.Println(x[i], y[i])
    }
}