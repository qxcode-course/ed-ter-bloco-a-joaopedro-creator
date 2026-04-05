package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct{
	l, c int
}

func search(grid [][]rune, atual Pos, endPos Pos) bool {
	if atual.l < 0 || atual.l >= len(grid[0]) || atual.c < 0 || atual.c >= len(grid[0]){
		return false
	}

	if grid[atual.l][atual.c] != ' ' {
		return false
	}

	grid[atual.l][atual.c] = '.'

	if atual.l == endPos.l && atual.c == endPos.c {
		return true
	}

	vizinhos := []Pos{
		{atual.l - 1, atual.c},
		{atual.l + 1, atual.c},
		{atual.l, atual.c - 1},
		{atual.l, atual.c + 1},
	}

	for _, v := range vizinhos {
		if search(grid, v, endPos){
			return true
		}
	}

	grid[atual.l][atual.c] = ' '
	return false


}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}


	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
