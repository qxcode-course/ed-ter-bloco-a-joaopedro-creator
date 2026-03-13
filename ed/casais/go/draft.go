package main
import "fmt"
func main() {
    var N int 
    fmt.Scan(&N)

    descasados := make(map[int]int)
    casados := 0

    for i := 0; i < N; i++ {
        var animal int
        fmt.Scan(&animal)

        if descasados[-animal] >0 {
            descasados[-animal]--
            casados ++
        } else {
            descasados[animal]++
        }
    }

    fmt.Println(casados)
}
