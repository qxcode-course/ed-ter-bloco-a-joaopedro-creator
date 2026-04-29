package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type  node struct {
    Value int            
    next *node          
    prev *node  
	root *node       
}

func (n *node) Next() *node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *node) Prev() *node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

type Llist struct{
	head * node
	size int
}

func NewLList() *Llist{
	lList := &Llist{}
	lList.head = &node{}
	lList.head.next = lList.head
	lList.head.prev = lList.head
	return lList
}


func (l *Llist) String() string {
	var elements []string
	for it := l.head.next; it != l.head; it = it.next {
		elements = append(elements, strconv.Itoa(it.Value))
	}
	return "[" + strings.Join(elements, ", ") + "]"
}

func (l *Llist) Size() int {
    return l.size
}

func (l *Llist) Insert(nodeRef *node, value int) {
	A := nodeRef
	B := nodeRef.prev
	C := &node{Value: value,
		next: A,
		prev: B,
		root: l.head,
	}

	B.next = C
	A.prev = C
	l.size++
}

func (l *Llist) PushBack(v int){
	l.Insert(l.head, v)
}
func (l *Llist) PushFront(v int){
	l.Insert(l.head.next, v)
}
func (l *Llist) PopBack(v int){
	l.Remove(l.head.prev)
}
func (l *Llist) PopFront(v int){
	l.Remove(l.head.next)
}

func (l *Llist) Clear(value int) {
	l.head.next = l.head
	l.head.prev = l.head
	l.size = 0
}


func (l *Llist) Remove(n *node) *node {
	if n == l.head {
		return nil
	}
	proximo := n.next
	n.prev.next = n.next
	n.next.prev = n.prev
	l.size--

	if proximo == l.head {
		return nil
	}
	return proximo
}

func (l *Llist) Front() *node {
	if l.size == 0 { return nil }
	return l.head.next
}

func (l *Llist) Back() *node {
	if l.size == 0 { return nil }
	return l.head.prev
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "show":
			fmt.Println(ll.String())
		case "size":
		    fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack(ll.head.Value)
		case "pop_front":
			ll.PopFront(ll.head.Value)
		case "clear":
			ll.Clear(ll.head.Value)
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	node.Value = newvalue
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "insert":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Insert(node, newvalue)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
