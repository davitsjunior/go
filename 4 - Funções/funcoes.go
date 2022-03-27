package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	fmt.Println(somar(1, 2))

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto função 01")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma2, _ := calculosMatematicos(20, 30)
	fmt.Println(resultadoSoma2)
}
