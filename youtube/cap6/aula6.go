// Fluxo de Controle
// Loops: utilizando ascii

package main

import "fmt"

func main() {

	for x := 33; x <= 122; x++ {
		//fmt.Printf("%#U\n", x)
		fmt.Printf("%d - %v \n", x, string(x))
	}

}
