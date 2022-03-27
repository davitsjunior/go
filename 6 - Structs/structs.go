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

	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome, u.idade = "Davi", 21
	fmt.Println(u)

	endereco := endereco{"Rua Pernambuco", 75}

	usuario2 := usuario{"Davi", 38, endereco}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)

}
