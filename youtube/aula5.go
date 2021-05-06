package main

import (
	"fmt"
)

//var x int = 10
var x int

//x = 15 - gera erro
func main() {

	x = 20
	fmt.Println(x)
	fmt.Printf("%v. %T \n", x, x)
}
