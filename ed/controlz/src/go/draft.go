package main
import (
    "fmt"
    "os"
    "bufio"
)
func main() {
     scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    entrada := scanner.Text()

     texto := []rune{}
    cursor := 0
    for _, ch := range entrada {
        switch ch{
    case 'R' :
        texto = append(texto[:cursor], append([]rune{'\n'}, texto[cursor:]...)...)
            cursor++
    case 'B':
        if cursor > 0 {
            texto = append(texto[:cursor-1], texto[cursor:]...)
            cursor--
        }
    
    case 'D' :
        if cursor < len(texto) {
            texto = append(texto[:cursor], texto[cursor+1:]...)
        }
    case '>' :
        if cursor < len(texto) {
            cursor++
        }
    case '<' :
        if cursor > 0 {
            cursor--
        }
    default: 
            texto = append(texto[:cursor], append([]rune{ch}, texto[cursor:]...)...)
            cursor++
    }
    }
    texto = append(texto[:cursor], append(texto[cursor:], []rune{'|'}...)...)
    fmt.Println(string(texto))
}