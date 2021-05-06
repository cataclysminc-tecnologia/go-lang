package main

import (
	"fmt"
)

var y = 111 // global

func main() {
	y := 666
	qualquercoisa(y)
}

func qualquercoisa(x int) {
	fmt.Println(x)
	fmt.Println(y) // undefined se não for global
}
