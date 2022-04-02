package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {

	u1 := usuario{"Davi", 37}
	u1.salvar()

	u2 := usuario{"Lavis", 8}
	maioridade := u2.maiorIdade()
	fmt.Println(maioridade)

	u3 := usuario{"Gabriel", 10}
	u3.aniversario()
	fmt.Println(u3)

}
