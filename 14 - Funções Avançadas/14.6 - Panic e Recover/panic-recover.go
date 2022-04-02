package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != (nil) {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func media(n1, n2 float64) bool {
	defer recuperarExecucao()
	m := (n1 + n2) / 2

	if m > 6 {
		return true
	} else if m < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {

	fmt.Println(media(6, 8))
	fmt.Println("Pós Execução")
}
