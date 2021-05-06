package main

import (
	"fmt"
)

const x = 10 // não é nada, somente será quando for usado, assim sendo, só é definido no momento do uso

var y = 10

//var c int

var d float64

func main() {

	d = x

	fmt.Println(d)

}
