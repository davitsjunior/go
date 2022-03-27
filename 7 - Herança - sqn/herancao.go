package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    //herança não especifica o tipo, ele passa tudo a struc pessoa para estudante
	curso     string
	faculdade string
}

func main() {

	fmt.Println("Herança só que não")

	p1 := pessoa{"João", "Predro", 38, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "ESSEI"}
	fmt.Println(e1)
}
