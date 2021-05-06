// Fluxo de Controle
// Loops: break & continue

package main

import "fmt"

func main() {

	for x := 0; x < 20; x++ {
		//if x%2 != 0 {
		if x == 3 {
			// faz isso quando o número não é par
			//continue
			break
		}
		fmt.Println(x)

	}

}
