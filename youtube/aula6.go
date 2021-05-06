// Valor zero

package main

import (
	"fmt"
)

var a int
var b float64
var c string
var d bool
var x int

func main() {

	x = 10
	fmt.Printf("%v, %T \n", x, x)

	y := 666
	fmt.Printf("%v, %T \n", y, y)

	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", d, d)

}
