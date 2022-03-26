package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 1000000 //int8, int16, int32, somente int assumira a arquitetura do computador, aceita + e -
	fmt.Println(numero)

	var numero2 uint32 = 2 // mesmo esquema do de cima, porém, somente números positivos
	fmt.Println(numero2)

	//Alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//números reais

	var numeroReal1 float32 = 123.45 //sempre ponto vírgula não funciona
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45 //sempre ponto vírgula não funciona
	fmt.Println(numeroReal2)

	//Strings

	var str string = "Texto"
	fmt.Println(str)

	//char retorna número da tabela asc (um caracter) tipo init
	char := 'j'
	fmt.Println(char)

	//todo tipo de dado tem um valor inicial string vazio numério 0 erro nil

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

	erro = errors.New("Erro interno")
	fmt.Println(erro)
}
