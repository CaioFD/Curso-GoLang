package main

import "fmt"

// Go e fortemente tipado
func main() {
	// Primeira maneira de declarar uma variavel
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)


	 // Segunda maneira de declarar uma variavel
	variavel2 := "Variavel 2" 
	fmt.Println(variavel2)


	// Terceira maneira de declara uma variavel
	var (
		variavel3 string = "aaaaaa"
		variavel4 string = "bbbbbb"
	)
	fmt.Println(variavel3, variavel4)

	// Quarta maneira de declarar uma variavel
	variavel5, variavel6 := "Variavel 5" , "Variavel 6"
	fmt.Println(variavel5,variavel6)

	// As declaracoes de variaveis servem tambem para declarar uma constante
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// ...

	//No GO conseguimos fazer a troca de valores de uma variavel para outra sem utilizar um auxiliar
	variavel5,variavel6 = variavel6,variavel5
	fmt.Println("Valor da variavel 5 = " ,variavel5, "\n", "Valor da variavel 6 = ",variavel6)
}