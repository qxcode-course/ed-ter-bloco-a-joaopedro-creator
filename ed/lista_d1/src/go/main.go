package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type  root struct {
    Value int            
    next *root           
    prev *root           
}

type Llist struct{
	head * root
	size int
}

func NewLList() *Llist{
	lList := &Llist{}
	lList.head = &root{}
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

func (l *Llist) insert(A *root, value int) {
	B := A.prev
	
	C := &root{Value: value, next: A, prev: B}

	B.next = C
	A.prev = C
	l.size++
}

func (l *Llist) remove(node *root) {
	if node == l.head {
		return
	}
	A := node.prev
	B := node.next

	A.next = B
	B.prev = A
	l.size--
}

func (l *Llist) Size() int {
    return l.size
}

func (l *Llist) PushBack(value int) {
	l.insert(l.head, value) 
}

func (l *Llist) PushFront(value int) {
	l.insert(l.head.next, value) 
}

func (l *Llist) PopFront(value int) {
	l.remove(l.head.next) 
}

func (l *Llist) PopBack(value int) {
	l.remove(l.head.prev) 
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
			// ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
