package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
    return &MultiSet{
        data:     make([]int, capacity),
        size:     0,
        capacity: capacity,
    }
}


func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (vec *MultiSet) String() string {
	return "[" + Join(vec.data[0:vec.size], ", ") + "]"
}

func (ms *MultiSet) Insert(value int) {
	if ms.capacity == ms.size{
		if ms.capacity == 0{
			ms.capacity = 1
		} else {
			
			novo:= NewMultiSet(ms.capacity*2)
			for i:= 0; i < ms.size; i++{
				novo.data[i] = ms.data[i]
			}

			ms.data = novo.data
			ms.capacity = novo.capacity

		}
	} 

	indice:= ms.size
	for i:= 0; i < ms.size; i++{
		if ms.data[i] >= value{
			indice = i
			break
		}
	}
	for i:= ms.size; i > indice; i--{
		ms.data[i] = ms.data[i-1]
	}
	ms.size++
	ms.data[indice] = value
}

func (vec *MultiSet) Erase(index int) error {
	if index < 0 || index >= vec.size {
        return fmt.Errorf("value not found")
    }

	for i := index; i < vec.size - 1; i++{
		vec.data[i] = vec.data[i + 1]
	}
	vec.size--
	return nil
}

func (vec *MultiSet) Contains(value int ) bool {
	for i := 0; i < vec.size; i++{
		if vec.data[i] == value{
			return true 
		}
	}

	return false
}

func (ms *MultiSet) Count(value int) int{
	total := 0
	for i:= 0; i < ms.size; i++{
		if value == ms.data[i] {
			total++
		}
	}
	return total
}

func (ms *MultiSet)Unique() int{
	total := 0
	num := 0
	for i := 0; i < ms.size; i++{
		if ms.data[i] > num{
			num = ms.data[i]
			total++
		}
	}
	return total
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			 value, _ := strconv.Atoi(args[1])
			 err := ms.Erase(value)
			 if err != nil{
				fmt.Println(err)
			 }
		case "contains":
			 value, _ := strconv.Atoi(args[1])
			 if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			 value, _ := strconv.Atoi(args[1])
			 fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
