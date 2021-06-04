package main

import "fmt"

func main() {
	var var1 int = 10
	//Essa linha baixo só faz uma cópia do valor da variável
	var var2 int = var1

	fmt.Println(var1, var2)

	//Isso vai mexer no valor apenas da variável 1
	var1++
	fmt.Println(var1, var2)

	//O ponteiro é uma referência de memória, no qual se alterar o valor da variável 1, se a 2 esiver apontando para o endereço da 1, vai printar o valor correto em ambas
	var var3 int = 100
	var pont *int = &var3

	//Isso vai printar o valor da var3 e o valor contido no endereço no pont
	fmt.Println(var3, *pont)

	//Isso vai printar o valor da var3 e o endereço de memória contido na variável pont
	fmt.Println(var3, pont)

	var3 = 200
	fmt.Println(var3, *pont)
}
