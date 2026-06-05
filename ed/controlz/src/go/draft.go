package main
import (
    "fmt"
    "os"
    "bufio"
)



type Estado struct {
	texto  []rune
	cursor int
}

func copia(texto []rune, cursor int) Estado {
	t := make([]rune, len(texto))
	copy(t, texto)
	return Estado{t, cursor}
}


func main() {
     scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    entrada := scanner.Text()

     texto := []rune{}
    cursor := 0

 
    var undoStack []Estado
    var redoStack []Estado
    
    for _, ch := range entrada {
    anterior := copia(texto, cursor)
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
    case 'Z' :
     if len(undoStack) > 0 {
            redoStack = append(redoStack, copia(texto, cursor))
            ultimo := undoStack[len(undoStack)-1]
            undoStack = undoStack[:len(undoStack)-1]
            texto, cursor = ultimo.texto, ultimo.cursor
        }
        
    case 'Y' :
         if len(redoStack) > 0 {
            undoStack = append(undoStack, copia(texto, cursor))
            ultimo := redoStack[len(redoStack)-1]
            redoStack = redoStack[:len(redoStack)-1]
            texto, cursor = ultimo.texto, ultimo.cursor
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
    if ch != 'Z' && ch != 'Y' && string(anterior.texto) != string(texto) {
        undoStack = append(undoStack, anterior)
        redoStack = nil
    }
    }
    texto = append(texto[:cursor], append([]rune{'|'}, texto[cursor:]...)...)
    fmt.Println(string(texto))
}