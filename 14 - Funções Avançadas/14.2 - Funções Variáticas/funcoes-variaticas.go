package main

import "fmt"

func soma(numeros ...int) int { //VARIÁTICO TEM QUE SER O ÚLTIMO PARÂMETRO E SÓ PODE HAVER UM (HIGHTLANDER)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func teste(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))

	teste("Número", 01, 02, 03, 04, 05, 06)
}
