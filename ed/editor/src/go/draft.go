package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    entrada := scanner.Text()

     texto := []rune{}
    cursor := 0

    for _, ch := range entrada{
    switch ch{
    case 'R' :
        texto = append(texto[:cursor], append([]rune{'\n'}, texto[cursor:]...)...)
            cursor++
    case 'B':
        
    
    case 'D' :
         
    case '>' :
          
    case '<' :
       
    default: // letras a-z e '-'
            
    }
    }
    texto = append(texto[:cursor], append([]rune{'|'}, texto[cursor:]...)...)
    fmt.Print(strings.ReplaceAll(string(texto), "\n", "\n"))
}
