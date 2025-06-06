package main

import "fmt"

func main() {
	//ARITIMÉTICOS
	soma := 1 + 2
	subracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	restoDivisao := 10 % 3

	fmt.Println(soma, subracao, divisao, multiplicacao, restoDivisao)

	//ATRIBUIÇÃO
	var variavel01 string = "String"
	variavel02 := "String 2"
	fmt.Println(variavel01, variavel02)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro || falso)

	//OPERADORES UNÁRIO
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 15
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero *= 15
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero /= 15
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
}
