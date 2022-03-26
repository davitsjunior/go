package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Gerenado Build")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("davijunior.com")
	fmt.Println(erro)
}
