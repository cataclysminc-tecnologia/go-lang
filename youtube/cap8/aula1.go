// Agrupamento de Dados - 1. Array

package main

import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", y)
	fmt.Println("Quantidade de itens do array x: ", len(x))
	fmt.Println("Quantidade de itens do array y: ", len(y))
}
