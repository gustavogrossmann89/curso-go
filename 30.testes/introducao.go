package main

import (
	"fmt"

	"introducao/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)

	tipoEndereco = enderecos.TipoDeEndereco("Av Paulista")
	fmt.Println(tipoEndereco)
}
