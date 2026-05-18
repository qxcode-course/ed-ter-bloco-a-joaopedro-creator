package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	count := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '1' {
			return
		}
		grid[i][j] = '0' 
		dfs(i-1, j)       
		dfs(i+1, j)       
		dfs(i, j-1)       
		dfs(i, j+1)       
	}
	 for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '1' {
                dfs(i, j) // afunda toda a ilha
                count++
            }
        }
    }

    return count
	
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
