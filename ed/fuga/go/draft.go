package main
import "fmt"
func main() {
    var H, P, F, D int

    fmt.Scan(&H, &P, &F, &D)

    for{
        F = (F + D + 16) % 16

        if F == H {
            fmt.Println("S")
            break
        }

        if F == P {
            fmt.Println("N")
            break
        }
    }
    
}
