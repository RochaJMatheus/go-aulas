package enderecos

import "strings"

//TipoDeEndereço verifica se um endereço tem um tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	primeiraPalavraDoEndereco := strings.Split(strings.ToLower(endereco), " ")[0]

	enderecoTemTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}
	if enderecoTemTipoValido {
		return primeiraPalavraDoEndereco
	}
	return "Tipo " + endereco + " inválido"
}
