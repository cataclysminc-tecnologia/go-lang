// Agrupamento de Dados - 2. Slice: literal composta

package main

import (
	"fmt"
)

func main() {
	slice := []string{"banana", "maçã", "jaca", "pêssego"}

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor: ", valor)
	}

	slice = append(slice, "melancia")

	for índice, valor := range slice {
		fmt.Println("No índice", índice, "temos o valor: ", valor)
	}

	for índice, _ := range slice {
		fmt.Println("Esse slice tem ", índice, " elementos.")
	}

	for _, valor := range slice {
		fmt.Printf("Um dos valores desse slice é %v. \n", valor)
	}

	slice2 := []int{20, 21, 22, 23}
	total := 0
	for _, valor := range slice2 {
		total += valor // mesma coisa que total = total + valor
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}

}
