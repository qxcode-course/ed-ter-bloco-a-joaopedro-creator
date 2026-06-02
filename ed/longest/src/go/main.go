package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	maior :=0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
			return 0
		}
		current := matrix[i][j]
		matrix[i][j] = -1
		maxPath := 1
		directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < len(matrix) && nj >= 0 && nj < len(matrix[0]) && matrix[ni][nj] > current {
				pathLength := 1 + dfs(ni, nj)
				if pathLength > maxPath {
					maxPath = pathLength
				}
	return 0
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
