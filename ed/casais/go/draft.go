package main

import "fmt"


func processarEntrada(restantes int, descasados map[int]int, casados int) int {
	// Caso base: não há mais animais para processar
	if restantes == 0 {
		return casados
	}

	var animal int
	fmt.Scan(&animal)

	// Lógica de negócio
	if descasados[-animal] > 0 {
		descasados[-animal]--
		casados++
	} else {
		descasados[animal]++
	}

	// Chamada recursiva para o próximo animal
	return processarEntrada(restantes-1, descasados, casados)
}

func main() {
	var N int
	if _, err := fmt.Scan(&N); err != nil {
		return
	}

	descasados := make(map[int]int)
	
	// Iniciamos a recursão passando N, o mapa vazio e 0 casados
	totalCasados := processarEntrada(N, descasados, 0)

	fmt.Println(totalCasados)
}




