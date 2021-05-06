// Conversão de tipos

package main

import (
	"fmt"
)

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%v\n", x)
	//fmt.Printf("%T \n", b) // hotdog

	x = int(b)
	fmt.Printf("%v \n", x)

}
