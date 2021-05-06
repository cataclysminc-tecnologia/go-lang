// Deslocamento de bits

package main

import (
	"fmt"
)

func main() {
	x := 24
	y := x << 1 // deslocando para a esquerda
	z := x >> 1 // deslçocando para a direita
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)

}
