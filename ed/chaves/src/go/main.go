package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	scanner := bufio.NewReader(os.Stdin)
	q := NewQueue[string]()
	for i := 0; i < 16; i++ {
		q.Enqueue(string(rune('A' + i)))
	}

	for q.items.Len() > 1{
		esquerda := q.Dequeue()
		direita := q.Dequeue()

		var Egols, Dgols int
		 fmt.Fscan(scanner, &Egols, &Dgols)

		if Egols > Dgols {
			q.Enqueue(esquerda)
		} else {
			q.Enqueue(direita)
		}
	}
	 fmt.Println(q.Dequeue())
}
