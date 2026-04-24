package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity), 
		size:     0,
		capacity: capacity,
	}
}

func (vec *Set) String() string {
	return "[" + Join(vec.data[0:vec.size], ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}


func (set *Set) Insert(value int) {
	for i := 0; i < set.size; i++{
		if set.data[i] == value{
			return
		}
	}
	if set.size == set.capacity{
		newCapacity := set.capacity * 2
		if newCapacity == 0 {
			newCapacity = 1
		}
		novo := NewSet(newCapacity)
		for i:= 0; i < set.capacity; i++{
			novo.data[i] = set.data[i]
		}
		set.data = novo.data
		set.capacity = novo.capacity
	}
	
	indice := set.size
	for i:= 0; i < set.size; i++{
		if value < set.data[i]{
			indice = i
			break
		}
	}
	for j := set.size; j > indice; j--{
			set.data[j] = set.data[j - 1]
		}
	set.data[indice] = value
	set.size++
	

	
}

func (vec *Set) Erase(value int) bool {
    indice := -1
    for i := 0; i < vec.size; i++ {
        if value == vec.data[i] {
            indice = i
            break
        }
        
        if vec.data[i] > value {
            break
        }
    }
   
    if indice == -1 {
        return false
    }

    for i := indice; i < vec.size-1; i++ {
        vec.data[i] = vec.data[i+1] 
    }

    vec.data[vec.size-1] = 0
    vec.size--
    return true
}

func (vec *Set) Contains(value int ) bool {
	for i := 0; i < vec.size; i++{
		if vec.data[i] == value{
			return true 
		}
	}

	return false
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	 v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			 for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
				v.Insert(value)
			 }
		case "show":
			fmt.Println(v)
		case "erase":
			 value, _ := strconv.Atoi(parts[1])
			 if !v.Erase(value){
				fmt.Println("value not found")
			 }

		case "contains":
			 value, _ := strconv.Atoi(parts[1])
			 fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
