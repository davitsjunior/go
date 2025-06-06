package main

import "fmt"

func main() {

	var variavel1 string = "Variavel 1"
	variavel2 := "Variável 2" //inferência de tipo

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string
		variavel4 string = "Teste 02"
	)

	fmt.Println(variavel3 + " " + variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel5, " ", variavel6)

	const constante1 string = "Constante 01" //segue as mesmas regras da declaração das variáveis
	fmt.Println(constante1)

}
