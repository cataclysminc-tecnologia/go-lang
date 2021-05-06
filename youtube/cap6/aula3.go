// Fluxo de Controle
// nested loop (repetição hierárquica)

package main

import "fmt"

func main() {

	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês: ", mes)
		for x := 0; x < 60; x++ {
			fmt.Print("Dia: ", x, ", ")
		}
		fmt.Println("\n")
	}

}
