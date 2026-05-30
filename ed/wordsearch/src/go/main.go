package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	for i := range grid {
		for j := range grid[i]{
			if dfs(grid, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(grid [][]byte, word string,lin, col, index int) bool {
if index == len(word) {
	return true
}
if lin < 0 || lin >= len(grid) || col < 0 || col >= len(grid[lin]) || grid[lin][col] != word[index] {
	return false
}
temp := grid[lin][col]
grid[lin][col] = '#'
if dfs(grid, word, lin+1, col, index+1) || dfs(grid, word, lin-1, col, index+1) || dfs(grid, word, lin, col+1, index+1) || dfs(grid, word, lin, col-1, index+1) {
	return true
}
grid[lin][col] = temp
return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
