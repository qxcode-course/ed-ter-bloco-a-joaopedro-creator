package main

import (
    "bufio"
    "fmt"
    "os"
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

    return false
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    var seq string
    fmt.Fscan(reader, &seq)

    var L int
    fmt.Fscan(reader, &L)

    s := []byte(seq)

    
    var dots []int
    for i, c := range s {
        if c == '.' {
            dots = append(dots, i)
        }
    }

    backtrack(s, dots, 0, L)
    fmt.Println(string(s))
}
