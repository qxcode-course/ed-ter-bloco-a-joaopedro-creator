package main
import "fmt"
import "math"

func main() {
    var a, b, c float64;
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

    sp := (a + b +c)/2;
    heron := math.Sqrt(sp* (sp - a) * (sp - b) * (sp - c));

    fmt.Printf("%.2f\n", heron)
}
