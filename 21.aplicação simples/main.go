package main

import (
	"linha-de-comando/app"
	"os"
)

//Uso:
// 		go run .\main.go ip --host amazon.com.br
//		go run .\main.go servidores --host amazon.com.br
func main() {
	aplicacao := app.Gerar()
	aplicacao.Run(os.Args)
}
