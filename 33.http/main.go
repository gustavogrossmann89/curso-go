package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Gustavo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando página de usuários!"))
}

func main() {

	//Acessar localhost:5000/home no browser
	http.HandleFunc("/home", home)

	//Acessar localhost:5000/usuarios no browser
	http.HandleFunc("/usuarios", usuarios)

	//Servidor startado, rodando na porta 5000
	log.Fatal(http.ListenAndServe(":5000", nil))
}
