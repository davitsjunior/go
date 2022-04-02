package main

import "fmt"

func f1() {
	fmt.Println("Função 01")
}

func f2() {
	fmt.Println("Função 02")
}

func media(n1, n2 float32) bool {
	if ((n1 + n2) / 2) >= 6 {
		return true
	}
	return false
}

func main() {

	defer f1() //DEFER = ADIAR	adia execução até o final do processo
	f2()
	fmt.Println(media(8, 6))

}
