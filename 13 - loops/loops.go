package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Inclementando I")
		//time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ { //variavel J somente acessível no scope (FOR{})
		fmt.Println("Inclementando J", j)
	}

	nomes := [3]string{"Davi", "Gabriel", "Lavínia"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, nome := range "Lavínia" {
		fmt.Println(indice, nome)
	}

	for indice, nome := range "Lavínia" {
		fmt.Println(indice, string(nome))
	}

	usuario := map[string]string{
		"nome":      "David",
		"sobrenome": "Tomas",
	}

	for chave, valor := range usuario { //RANGE SÓ FUNCIONA EM MAP NÃO EM STRUCTS
		fmt.Println(chave, valor)
	}

}
