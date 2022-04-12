package main

import (
	"fmt"
	"time"
)

func main() {

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		canal <- fmt.Sprintf("Valor recebido: %s", texto)
		time.Sleep(time.Millisecond * 500)
	}()

	return canal
}
