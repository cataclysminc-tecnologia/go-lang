// Criando meus próprios tipos

package main

import (
	"fmt"
)

type hotdog int

var b hotdog

func main() {
	x := 10

	fmt.Printf("%T \n", x)
	fmt.Printf("%T \n", b) // hotdog

	b = x
}
