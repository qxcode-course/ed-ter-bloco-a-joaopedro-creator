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

func burnTrees(grid [][]rune, l, c int) {

	if !is_value(grid, l, c, '#') {
		return
	}

	
	grid[l][c] = 'o'

	
	burnTrees(grid, l, c+1)
	burnTrees(grid, l, c-1)
	burnTrees(grid, l+1, c)
	burnTrees(grid, l-1, c)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for i := 0; i < nl; i++ {
		if scanner.Scan() {
			line := []rune(scanner.Text())
			grid = append(grid, line)
		}
	}

	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}