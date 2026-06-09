package main

import (
	//"bufio"
	"fmt"
	//"os"
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
	var linhas, colunas int
	fmt.Scanf("%d %d", &linhas, &colunas)	

	grid := make([][]rune, linhas)
    var start, end Pos

    for i := 0; i < linhas; i++ {
        var line string
        fmt.Scan(&line) 
        grid[i] = []rune(line)
    }

	for l := 0; l < linhas; l++ {
		for c := 0; c < colunas; c++ {
			if grid[l][c] == 'S' {
				start = Pos{l, c}
			} else if grid[l][c] == 'F' {
				end = Pos{l, c}
			}
		}
	}

	dfs(grid, start, end)
}


