package main

import (
	"errors"
	"fmt"
)

// Demonstração de tipos de dados
// int8, int16, int32, int64 : int com determinada quantidade de bits
// uint : unsigned int
// uint8, uint16, uint32, uint64
func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//rune = int32
	var numero3 rune = 1000000
	fmt.Println(numero3)

	//alias
	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numReal1 float32 = 123.45
	var numReal2 float64 = 123.45

	fmt.Println(numReal1)
	fmt.Println(numReal2)

	//STRINGS
	var str string = "String de teste"
	fmt.Println(str)

	//Praticamente não existe CHAR
	//Se fizer assim com aspas simples, vai devolver o int da tabela ASCII
	char := 'B'
	fmt.Println(char)

	//Não declarar valor para variável, não vai dar erro
	//Todo tipo de dado tem um valor inicial padrão: "", 0, nil

	//Tipo booleano
	var booleano1 bool = true
	var booleano2 bool = false

	fmt.Println(booleano1)
	fmt.Println(booleano2)

	//Tipo error
	var erro error = errors.New("Novo erro")
	fmt.Println(erro)
}
