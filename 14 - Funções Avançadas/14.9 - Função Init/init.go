package main

import "fmt"

func init() {
	fmt.Println("Print Init") //incicializa antes do main (primeiro função a executar, não importa se está no comeõ ou no fim do arquivo)
}

func main() {
	fmt.Println("Pinto Main")
}
