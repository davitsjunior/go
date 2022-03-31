package main

import "fmt"

func main() {
	fmt.Println("Estutura de Controle")

	n1 := 10

	if n1 > 15 {
		fmt.Println("Maior que quinze")
	} else {
		fmt.Println("Menor ou igual a quinze")
	}

	if n2 := n1; n2 > 0 { //variavel usando ifinit ela somente funcionará no scopo (IF {})
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Número é menor que zero")
	}

}
