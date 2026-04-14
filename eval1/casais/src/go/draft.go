package main
import "fmt"

func entrada(restantes int, descasados map[int]int, casados int)int{
    if restantes == 0{
        return casados
    }
    var animal int
    fmt.Scan(&animal)

    if descasados[-animal] > 0 {
        descasados[-animal]--
        casados++
    }else{
        descasados[animal]++
    }

    return entrada(restantes - 1, descasados, casados)
}
func main() {
    var n int 
    fmt.Scan(&n)

    descasados := make(map[int]int)
    totalCasados := entrada(n, descasados, 0)

    fmt.Println(totalCasados)



}
