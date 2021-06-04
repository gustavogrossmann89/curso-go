package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "Nome1"
	fmt.Println(array2)

	array3 := [5]string{"Nome1", "Nome2", "Nome3", "Nome4", "Nome5"}
	fmt.Println(array3)

	//Na prática, esse exemplo abaixo não é muito utilizado
	//Até o array no geral não é, por ser fixo
	array4 := [...]string{"Nome1", "Nome2", "Nome3"}
	fmt.Println(array4)

	//Já o slice é bem mais utilizado, por ser dinâmico
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)

	//Adicionar item no slice
	slice = append(slice, 10)
	fmt.Println(slice)

	//Esse caso abaixo, funciona como um ponteiro
	slice2 := array3[1:3]
	fmt.Println(slice2)

	//Se alterar o valor do array3, vai alterar o slice, pois o slice funciona como um ponteiro
	array3[2] = "NomeAlterado"
	fmt.Println(slice2)

	//Verificando o tipo de cada variável
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array4))

	//Segundo parâmetro é o tamanho e o terceiro é a capacidade máxima
	slice3 := make([]float32, 3, 3)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //length
	fmt.Println(cap(slice3)) //capacidade
	//O make acima cria um array de 10 posições e faz um slice dele de 3

	//Ao adicionar mais do que a capacidade permite, o GO vai duplicar o tamanho da capacidade
	slice3 = append(slice3, 20)
	fmt.Println(slice3)
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 20)
	fmt.Println(slice3)
	fmt.Println(cap(slice3)) //capacidade
}
