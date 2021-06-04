package main

//Import com um _ na frente, é um import implícito
//Nesse caso, tem que ser assim, pois se não, ao salvar, o GO remove o pacote
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Foi feito o go get github.com/go-sql-driver/mysql para baixar a dependência
func main() {
	stringConexao := "golang:Golang_2021$@/tutorial?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}
	//Mandando fechar a conexão no final
	defer db.Close()

	//Se mudarmos o usuário, senha ou base, com algum dado inválido, vai cair nesse erro abaixo
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão está aberta")

	//Consultando dados de uma tabela
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	//Sempre bom fechar antes de terminar
	defer linhas.Close()
}
