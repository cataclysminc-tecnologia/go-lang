package main

import (
	"fmt"
)

func main() {

	// Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Operadores
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Println(true && true, verdadeiro, falso)

	// Opeadores Unários
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Mneor que 5"
	}
	fmt.Println(texto)

}
