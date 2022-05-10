package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("avenida Paulista")

	fmt.Println(tipoEndereco)
}
