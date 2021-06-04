package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoMinuscula, " ")[0]

	isValido := false

	for _, tipo := range tiposValidos {
		if primeiraPalavraEndereco == tipo {
			isValido = true
			break
		}
	}

	if isValido {
		return strings.Title(primeiraPalavraEndereco)
	}

	return "Tipo inválido!"
}
