package main

import "fmt"

func clousure() func() {
	texto := "Print Clousure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {

	texto := "Print main"
	fmt.Println(texto)

	funcaoNova := clousure()
	funcaoNova()

}
