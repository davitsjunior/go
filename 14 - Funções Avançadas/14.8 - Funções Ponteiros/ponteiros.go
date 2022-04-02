package main

import "fmt"

func inverterSinal(n1 int) int {
	return n1 * -1
}

func inverterSinalPonteiro(n1 *int) {
	*n1 = *n1 * -1
}

func main() {

	numero := 20
	numeroinvertido := inverterSinal(numero)
	fmt.Println(numeroinvertido)
	fmt.Println(numero)

	numeroPonteiro := 40
	fmt.Println(numeroPonteiro)
	inverterSinalPonteiro(&numeroPonteiro)
	fmt.Println(numeroPonteiro)
}
