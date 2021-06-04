package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//Esse Template não é um slice, mas sim uma estrutura
var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	//Acessar localhost:5000/home no browser
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		//Enviando dados dinâmicos
		//No html, basta chamar com {{ .Propriedade }}
		u := usuario{"Gustavo", "gustavo.grossmann@gmail.com"}
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Listening on port 5000")
	//Servidor startado, rodando na porta 5000
	log.Fatal(http.ListenAndServe(":5000", nil))
}
