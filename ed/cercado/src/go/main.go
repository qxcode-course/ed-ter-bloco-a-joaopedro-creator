package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if board[i][j] != 'O'{
		return
	}
	board[i][j] = 'S'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}


// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	nl := len(board)
	nc := len(board[0])

	for i := 0; i < nl; i++ {
		dfs(board, i, 0)
		dfs(board, i, nc-1)
	}
	for j := 0; j < nc; j++ {
		dfs(board, 0, j)
		dfs(board, nl-1, j)
	}

	for i := 0; i < nl; i++ {
		for j := 0; j < nc; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'S' {
				board[i][j] = 'O'
			}
		}
	}

}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
