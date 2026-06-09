package main
import "fmt"

func gasEst(gas, dist []int) int {
    tanque, total, candidato := 0, 0, 0

    for i := 0; i < len(gas); i++{
        diff := gas[i] - dist[i]
		tanque += diff
		total += diff

        if tanque < 0{
            candidato = i + 1
            tanque = 0 
        }
    }

   if total < 0{
        return -1
   }    

    return candidato    

}

func main() {
    var n int 
    fmt.Scan(&n)

    gas := make([]int, n)
    dist := make([]int, n)

    for i := 0; i < n; i++ {
    fmt.Scan(&gas[i], &dist[i])
}
    
    fmt.Println(gasEst(gas, dist))  
}