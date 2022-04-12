package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá mundo"
	canal <- "Programando Go"

	mensagem01 := <-canal
	mensagem02 := <-canal

	fmt.Println(mensagem01)
	fmt.Println(mensagem02)
}
