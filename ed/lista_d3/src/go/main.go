package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}

func equals(a, b *LList) bool {
	if a.size != b.size {
		return false
	}
	na := a.root.next
	nb := b.root.next
	for na != a.root {
		if na.Value != nb.Value {
			return false
		}
		na = na.next
		nb = nb.next
	}
	return true
}

func addsorted(lla *LList, value int) {
	n := lla.root.next
	for n != lla.root && n.Value < value {
		n = n.next
	}
	lla.insertBefore(n, value)
}	

func reverse(lla *LList) {
	n := lla.root
	for {
		n.next, n.prev = n.prev, n.next
		n = n.prev 
		if n == lla.root {
			break
		}
	}
}
func merge(lla, llb *LList) *LList {
	merged := NewLList()
	na := lla.root.next
	nb := llb.root.next
	for na != lla.root && nb != llb.root {
		if na.Value < nb.Value {
			merged.PushBack(na.Value)
			na = na.next
		} else {
			merged.PushBack(nb.Value)
			nb = nb.next
		}
	}
	for na != lla.root {
		merged.PushBack(na.Value)
		na = na.next
	}
	for nb != llb.root {
		merged.PushBack(nb.Value)
		nb = nb.next
	}
	return merged
}


func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := merge(lla, llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
