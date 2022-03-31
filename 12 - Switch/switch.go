package main

import "fmt"

func diaDaSemana(numero int8) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido!"
	}
}

func diaDaSemana2(numero int8) string {
	switch {
	case numero == 1:
		return "Domingo"
	default:
		return "Somente teste"
	}
}

func main() {
	fmt.Println("Switch")
	dia := diaDaSemana(30)
	fmt.Println(dia)
}
