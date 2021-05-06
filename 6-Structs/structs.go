// Struct: Coleção de campos, cada campo tem um nome e um tipo

package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u) // Retorna { 0}

	u.nome = "Franciny"
	u.idade = 38
	fmt.Println(u)

	enderecoExemplo := endereco{"Travessa Juquia", 136}

	usuario2 := usuario{"Patty", 40, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Andres"}
	fmt.Println(usuario3)

}
