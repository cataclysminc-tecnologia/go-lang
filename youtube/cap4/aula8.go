/* IOTAS: It is an identifier which is used with constand
and which can simplify constant definitions that use auto increment numbers.
The IOTA keyword represent integer constant starting from zero. So essentially
it can be used to create effective constant in Go. They can also be used to create enum
in Go.
*/

package main

import (
	"fmt"
)

const (
	a = iota * 10000000000000000
	b
	c
	x = iota // 0
	_ = iota // 1
	z = iota // 2
)

func main() {

	fmt.Println(a, b, c, x, z)

}
