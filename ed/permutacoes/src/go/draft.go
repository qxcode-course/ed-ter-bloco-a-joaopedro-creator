package main
import (
    "fmt"
    // "strings"
    // "strconv"
)

func permut (prefixo string, restante string) {
    if len(restante) == 0 {
        fmt.Println(prefixo)
    } else {
        for i := 0; i < len(restante); i++ {
            permut(prefixo + string(restante[i]), restante[:i] + restante[i+1:])
        }
    }
}


func main() {
    permut( "", "abc")
}       
