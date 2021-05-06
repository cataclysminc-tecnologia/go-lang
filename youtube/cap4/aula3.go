package main

import (
	"fmt"
	"runtime"
)

var x bool
var i = 10.0

func main() {
	a := "e"
	b := "é"
	c := "是"
	fmt.Printf("%v, %v, %v \n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v, %v, %v \n", d, e, f)

	g := 10
	h := 10.0
	fmt.Printf("%v, %T \n", g, g)
	fmt.Printf("%v, %T \n", h, h) // float64
	fmt.Printf("%v, %T \n", i, i)

	fmt.Println(runtime.GOOS)   // SO
	fmt.Println(runtime.GOARCH) // Processador

}
