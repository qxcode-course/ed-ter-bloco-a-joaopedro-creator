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

func backtrack(s []byte, dots []int, idx, L int) bool {
   
    if idx == len(dots) {
        return true
    }

    pos := dots[idx]

    for d := byte('0'); d <= byte('0'+L); d++ {
        if isValid(s, pos, d, L) {
            s[pos] = d                                  
            if backtrack(s, dots, idx+1, L) {
                return true                             
            }
            s[pos] = '.'                               
        }
    }


func main() {
    
    entrada := bufio.NewScanner(os.Stdin)
    var L int
    var s string

    for entrada.Scan() {
        fmt.Sscanf(entrada.Text(), "%d %s", &L, &s)
        sBytes := []byte(s)
        var dots []int

        for i, c := range sBytes {
            if c == '.' {
                dots = append(dots, i)
            }
        }

        if backtrack(sBytes, dots, 0, L) {
            fmt.Println(string(sBytes))
        } else {
            fmt.Println("NO SOLUTION")
        }
    }
}
