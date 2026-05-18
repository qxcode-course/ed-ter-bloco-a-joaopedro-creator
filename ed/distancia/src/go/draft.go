package main
import (
    "fmt"
    "os"
    "bufio"
)


func isValid(s []byte, pos int, digit byte, L int) bool {
    for i, c := range s {
        if c == digit && i != pos {
            diff := pos - i
            if diff < 0 {
                diff = -diff
            }
            if diff < L {
                return false
            }
        }
    }
    return true
}


func main() {
    
}
