// Parecido ao struct

package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Franciny",
		"sobrenome": "Salles",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Andres",
			"ultimo":   "Rojas",
		},
		"curso": {
			"nome":   "Matemática",
			"campus": "Campus 666",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "escorpião",
	}

	fmt.Println(usuario2)
}
