// Agrupamento de Dados - 2. Slice: literal composta

package main

import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {
	array := [5]int{1, 2, 3, 4, 5} // Um array tem um numero finito de elementos especificado
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5} // um slice não tem um numero finito de elementos especificado
	fmt.Println(slice)

	//array2 := append(array, 6)
	slice2 := append(slice, 6) // Conseguimos mudar o tamanho do slice
	//fmt.Println(array2)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])
	slice[20] = 1
	fmt.Println(slice[20]) // gera erro: index out of range [20]
}
