package enderecos

import "strings"

//TipoDeEndereco Função exportada precisa começar com letra maiúscula
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]

	enderecoMinusculo := strings.ToLower(primeiraPalavraDoEndereco)

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == enderecoMinusculo {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return strings.Title(enderecoMinusculo)
	}

	return "Tipo Inválido"
}
