package main

import (
	"fmt"
)

func main() {
	//s := "Hello, playground"
	s := "ascii, éóâ, 文"
	sb := []byte(s)

	/*
		for _, v := range sb {
			fmt.Printf("%v - %T - %#U - %#x \n", v, v, v, v)
		}*/

	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x \n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x \n", s[i], s[i], s[i], s[i]) // byte por byte
	}

	fmt.Printf("%v \n %T", sb, sb)

}
