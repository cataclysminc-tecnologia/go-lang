package main

import (
	"fmt"
)

var a = "Bom dia" // variável global

func main() {

	// Declarando e atribuindo valor ao mesmo tempo
	x := 10
	y := "Good morning!"
	b := 10 + 10
	c := 10 == 10 // verifica se é igual - true
	d := 10 < 10  // false

	fmt.Println(x, y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x = 20 // atribuição
	x, z := 21, 30

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

	fmt.Printf("a: %v, %T\n", a, a)

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
