package main

import (
	"fmt"
	"sort"
)

// Função recursiva para ler as figurinhas e marcar as presentes
func lerFigurinhas(n, idx int, tem []bool, fig []int) {
	if idx == n {
		return
	}
	var v int
	fmt.Scan(&v)
	fig[idx] = v
	if v < len(tem) {
		tem[v] = true
	}
	lerFigurinhas(n, idx+1, tem, fig)
}

// Função recursiva para encontrar repetidas (exige slice ordenado)
func acharRepetidas(fig []int, idx int, acc []int) []int {
	if idx >= len(fig) {
		return acc
	}
	if fig[idx] == fig[idx-1] {
		// Evita adicionar o mesmo número repetido várias vezes se houver triplicatas
		if len(acc) == 0 || acc[len(acc)-1] != fig[idx] {
			acc = append(acc, fig[idx])
		}
	}
	return acharRepetidas(fig, idx+1, acc)
}

// Função recursiva para encontrar figurinhas faltando
func acharFaltando(tem []bool, idx int, acc []int) []int {
	if idx >= len(tem) {
		return acc
	}
	if !tem[idx] {
		acc = append(acc, idx)
	}
	return acharFaltando(tem, idx+1, acc)
}

// Helper para imprimir os slices conforme o formato solicitado
func imprimirResultado(lista []int) {
	if len(lista) == 0 {
		fmt.Println("N")
		return
	}
	for i, v := range lista {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	var total, n int
	if _, err := fmt.Scan(&total, &n); err != nil {
		return
	}

	fig := make([]int, n)
	tem := make([]bool, total+1)

	// 1. Leitura Recursiva
	lerFigurinhas(n, 0, tem, fig)

	// 2. Processamento de Repetidas
	// Ordenamos para que os iguais fiquem adjacentes, facilitando a lógica recursiva
	sort.Ints(fig)
	repetidas := acharRepetidas(fig, 1, []int{})
	imprimirResultado(repetidas)

	// 3. Processamento de Faltantes
	faltando := acharFaltando(tem, 1, []int{})
	imprimirResultado(faltando)
}