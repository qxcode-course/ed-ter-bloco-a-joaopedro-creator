package main

import (
	"bufio"
	"fmt"
	"os"
)

func is_value(grid [][]rune, l, c int, value rune) bool {
	nl := len(grid)
	nc := len(grid[0])

	
	if l < 0 || l >= nl || c < 0 || c >= nc {
		return false
	}

	return grid[l][c] == value
}

type Pos struct{ l, c int }
func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	
stack.Push(Pos{l, c})

	for !stack.IsEmpty() {
		pos := stack.Pop()
		if !is_value(grid, pos.l, pos.c, 'o') {
			continue
		}

		
		grid[pos.l][pos.c] = '#'

		
		stack.Push(Pos{pos.l, pos.c + 1})
		stack.Push(Pos{pos.l, pos.c - 1})
		stack.Push(Pos{pos.l + 1, pos.c})
		stack.Push(Pos{pos.l - 1, pos.c})
	}
	

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}




// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha