package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m , n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	result := 0
	for i := 0; i<m;i++{
		for j := 0; j<n;j++{
			result = max(result, dfs(matrix, i, j, memo))
		}
	}
	return result
}

func dfs(matrix [][]int, i, j int, memo [][]int) int {
	if memo[i][j] !=0{
		return memo[i][j]
	}
	m, n := len(matrix), len(matrix[0])
	maxResult :=1
	for _, d := range dirs{
		ni, nj := i + d[0], j + d[1]
		if ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j]{
			maxResult = max(maxResult, 1 + dfs(matrix, ni, nj, memo))
		}
	}
	memo[i][j] = maxResult

	return maxResult
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
