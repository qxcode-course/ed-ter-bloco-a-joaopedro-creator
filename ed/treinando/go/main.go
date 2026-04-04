package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func auxTostr(vet []int) string {
	if len(vet) == 0 {
		return ""
	}
	atual := strconv.Itoa(vet[0])
	if len(vet) == 1 {
		return atual
	}
	return atual + ", " + auxTostr(vet[1:])
}

func tostr(vet []int) string {
	return "[" + auxTostr(vet) + "]"
}

func auxTostrrev(vet []int) string {
	if len(vet) == 0 {
		return ""
	}
	atual := strconv.Itoa(vet[0])
	if len(vet) == 1 {
		return atual
	}
	return auxTostrrev(vet[1:]) + ", " + atual
}

func tostrrev(vet []int) string {
	return "[" + auxTostrrev(vet) + "]"
}

// reverse: inverte os elementos do slice
func auxReverse(vet []int, i int, j int) {
	if i >= j {
		return
	}
	vet[i], vet[j] = vet[j], vet[i]
	auxReverse(vet, i+1, j-1)
}

func reverse(vet []int) {
	if len(vet) > 0 {
		auxReverse(vet, 0, len(vet)-1)
	}
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 { 
		return -1 
	} 
	if len(vet) == 1 { 
		return 0 
	}  
		
	iResto := min(vet[1:]) + 1 

	if vet[0] <= vet[iResto] {
		return 0
	}
	return iResto
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}