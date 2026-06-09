package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct{
	l, c int
}
func dfs(grid [][]rune, start, end Pos) {
    caminho := NewStack[Pos]()
    becos := NewStack[Pos]()
    caminho.Push(start)

    for !caminho.IsEmpty() {
        atual := caminho.Top()
        grid[atual.l][atual.c] = '.'

        if atual == end {
            break
        }

        vizinhos := []Pos{
            {atual.l - 1, atual.c},
            {atual.l + 1, atual.c},
            {atual.l, atual.c - 1},
            {atual.l, atual.c + 1},
        }

        validos := []Pos{}
        for _, v := range vizinhos {
            if grid[v.l][v.c] == ' ' || grid[v.l][v.c] == 'F' {
                validos = append(validos, v)
            }
        }

        if len(validos) > 0 {
            caminho.Push(validos[0])
        } else {
            becos.Push(caminho.Pop())
        }
    }

    for !becos.IsEmpty() {
        b := becos.Pop()
        grid[b.l][b.c] = ' '
    }
}

func main() {
entrada := bufio.NewScanner(os.Stdin)
entrada.Scan()
var nl, nc int
fmt.Sscanf(entrada.Text(), "%d %d", &nl, &nc)
}


