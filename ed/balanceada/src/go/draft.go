package main
import "fmt"


func verificar(s string) bool{
    pilha := [] rune{}

    for _, c := range s{
        if c == '(' || c == '['{
            pilha = append(pilha, c)
        }else {
            if len(pilha) == 0{
                return false 
            }

            topo := pilha[len(pilha)-1]
            pilha = pilha[:len(pilha)-1]

        if c == ')' && topo != '('{
            return false
        }

        if c == ']' && topo != '['{
            return false
        }
        }

    
    }
     return len(pilha) == 0
}

func main() {
    var s string 
    fmt.Scan(&s)

    if verificar(s){
        fmt.Println("balanceado")
    }else{
        fmt.Println("nao balanceado")
    }
  
}