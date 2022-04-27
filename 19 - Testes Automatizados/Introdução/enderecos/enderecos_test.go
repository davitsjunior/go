package enderecos

import "testing"

type cenarioDeTeste struct{
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T){
	//TestXxxxXxXxxxxx

	cenariosDeTestes := [] cenarioDeTeste{
		{"Rua Pará", "Rua"},
		{"Avenida Pará", "Avenida"},
		{"Beco Pará", "Tipo Inválido"},
		{"Praça Pará", "Tipo Inválido"},
		{"Vila Pará", "Tipo Inválido"},
		{"Rodovia Pará", "Rodovia"},
		{"Viela Pará", "Tipo Inválido"},
		{"RUA Pará", "Rua"},
		{"Estrada Pará", "Estrada"},
	}

	for _, cenario := range cenariosDeTestes{
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.enderecoEsperado{
			t.Errorf("O Tipo de Endereço %s é diferente do esperado %s",
				tipoEnderecoRecebido,
				cenario.enderecoEsperado,
			)
		}
	}
}