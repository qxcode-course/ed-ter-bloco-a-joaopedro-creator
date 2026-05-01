package main

import (
	"fmt"
	//"strings"
)


// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	res := "[ "

	for n := l.Front(); n != l.End(); n = n.next{
		res += fmt.Sprint(n.Value)
		
		if n == sword{
			res += ">"
		}

		if n.next != l.End(){
			res += " "
		}
	}
	res += " ]" 
    return res
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	prox := it.next

	if prox == l.root{
		prox = prox.next
	}
	
	return prox
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	//fmt.Println(qtd, chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
