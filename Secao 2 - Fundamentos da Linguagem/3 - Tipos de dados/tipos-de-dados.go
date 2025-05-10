package main

import (
	"errors"
	"fmt"
)

func main() {
	// comecando pelos numeros:

	// 1 - INTEIROS:
	//int8, int16, int32, int64  a diferenca é a quantidade de bits que ele suporta, quanto maior o int maior o numero que ele ira suportar!
	fmt.Println("Tipos de INTEIROS em GO: ")
	var numero1 int8 = 100 // caso ultrapasse o numero de bits ele nao compilara por ter tido um overflow
	fmt.Println("int8: ", numero1)

	var numero2 int16 = 10000
	fmt.Println("int16: ", numero2)

	var numero3 int32 = 1000000000
	fmt.Println("int32: ", numero3)

	var numero4 int64 = 1000000000000000000
	fmt.Println("int64: ", numero4)

	var numero5 int = 1000000000000000000 // o int sozinho ira usar a arquitetura do seu computador como base, se ele for 64 bits o int sozinho sera igual a um int64
	fmt.Println("int SOLO: ", numero5)

	numero6 := 1000000000000000000 // Quando nao declarar o tipo explicitamente ele vai utilizar a inferencia de tipos e sera do tipo int e ele sera do tipo da arquitetura do computador
	fmt.Println("int_INFENCIA_DE_TIPO:",numero6)

	//podemos utilizar rune no lugar de int32, pois na documentacao do GO eles sao praticamente a mesma coisa. Chamado de alias.
	//usado principalemte para a tabela asc
	// FIM INTEIROS

	fmt.Println("\n")

	// 2 - REAIS:
	// float32, float64 segue a mesma logica de bits do int
	fmt.Println("Tipos de REAIS em GO: ")
	var numeroReal1 float32 = 1.234 
	fmt.Println("float32: ",numeroReal1, "...")

	var numeroReal2 float64 = 12.289371379
	fmt.Println("float64: ",numeroReal2, "...")

	numeroReal3 := 13.124 // a inferencia de tipo funciona para float quando voce coloca um numero quebrado, se for um numero "nao quebrado" ele sera declarado como int
	fmt.Println("float_INFENCIA_DE_TIPO: ", numeroReal3, "...") 

	fmt.Println("\n")
	// FIM REAIS

	// 3 - STRINGS:
	fmt.Println("Tipos de STRING em GO: ")
	
	var str string = "Texto" // cadei de caracteres
	fmt.Println("string: ", str)

	str2 := "Texto2"
	fmt.Println("string_INFENCIA_DE_TIPO: ",str2)

	// no Go "nao existe" o tipo CHAR
	char := 'A' // o char printa o numero que o caracter aparece na tabela ASC -> A = 65
	fmt.Println("Char da letra A: ", char)
	// FIM STRINGS

		fmt.Println("\n")

	// 4 - Booleano:
	fmt.Println("Tipos de BOOLEANO em GO: ")
		var boolean01 bool = true
		fmt.Println("Booleano 1: ", boolean01)

		var boolean02 bool = false
		fmt.Println("Booleano 2: ", boolean02)
	// FIM Booleano

	fmt.Println("\n")

	// 5 - ERRO!
	var erro error  = errors.New("Erro interno")  // para declara um valor ao erro vc deve utilizar um pacote do GO chamado "errors"
	fmt.Println("Erro: ", erro)

	fmt.Println("\n")

	// * 6 - Valores Nulos:
	fmt.Println("Tipos de VALORES NULOS em GO: ")
	// valores nulos sao chamados de valores iniciais ou valores zero(0)
	// Ocorre quando voce dclara ela com a clausula vaga e com o tipo de dado, mas nao coloca um valor

	var valorNULO1 string // Quando voce nao declara o valor na string ela retornara vazio
	fmt.Println("Valor Nulo para string: ", valorNULO1)

	var valorNULO2 int // Quando voce nao declara o valor no int ele retornara zero(0)
	fmt.Println("Valor Nulo para inteiro: ", valorNULO2)

	var valorNULO3 float64 // Quando voce nao declara o valor no float ele retornara zero(0)
	fmt.Println("Valor Nulo para float: ", valorNULO3)

	var valorNULO4 bool // Quando voce nao declara o valor no boolean ele retornara o valor de FALSE
	fmt.Println("Valor Nulo para boolean: ", valorNULO4)

	var valorNulo5 error // Quando voce nao declara o valor no erro ele retornara <nil>
	fmt.Println("Valor Nulo para erro: ", valorNulo5)
	// FIM Valores Nulos
	
}