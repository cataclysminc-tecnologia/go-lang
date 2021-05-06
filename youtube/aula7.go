// O pacote fmt

package main

import (
	"fmt"
)

func main() {

	a := "oi"
	x := "oi bom dia\n como vai? \t espero que tudo bem."
	y := `"oi bom dia\n como vai? \t espero que tudo bem."` // literal
	z := fmt.Sprint(x, " ", y)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Print(a)
	fmt.Println(z)

}
