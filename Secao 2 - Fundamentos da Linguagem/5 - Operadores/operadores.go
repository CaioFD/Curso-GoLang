package main

import "fmt"

func main() {
	// 1 - ARITMETICOS: + - / * %
	fmt.Println("-----------")
	fmt.Println("ARITMETICOS: ")
	soma := 2 + 2
	fmt.Println("Soma = ", soma)
	
	subtracao := 5 - 2
	fmt.Println("Subtracao = ", subtracao)

	multiplicacao := 2 * 7
	fmt.Println("Multiplicacao = ", multiplicacao)

	divisao := 10 / 2
	fmt.Println("Divisao = ", divisao)

	restodaDivisao := 10 % 2
	fmt.Println("Resto da Divisao = ", restodaDivisao)

// No GO voce nao consegue fazer contas aritmeticas com numeros de diferentes tipagens como:
// var numero1 int16 = 10
// var numero2 int32 = 10
// soma := numero1 + numero2
// fmt.Println(soma)

// FIM DOS ARITMETICOS

// 2 - ATRIBUICAO
fmt.Println("-----------")
fmt.Println("ATRIBUICAO: ")
	var VARIAVEL1 string = "String_Teste"
	VARIAVEL2 := "String_Teste2"
	fmt.Println( VARIAVEL1,VARIAVEL2)

// FIM DAS ATRIBUICOES

// 3 - RELACIONAIS
// operadores relacionais sempre retornam boolean (true or false)
fmt.Println("-----------")
fmt.Println("RELACIONAIS: ")
fmt.Println(1 > 2) 
fmt.Println(3 < 2)
fmt.Println(2 <= 2) 
fmt.Println(2 >= 2) 
fmt.Println(2 == 2)
fmt.Println(2 != 2) 
// FIM DOS RELACIONAIS

// 4 - LOGICOS
// Nos operadores logicos temos apenas 3: and, or, not
fmt.Println("-----------")
fmt.Println( "LOGICOS: ")
verdadeiro, falso := true,false
fmt.Println(verdadeiro && verdadeiro) // && -> and
fmt.Println(verdadeiro && falso) 

fmt.Println(verdadeiro || falso) // || -> or
fmt.Println(verdadeiro || verdadeiro) 
fmt.Println(falso || falso) 

fmt.Println(!verdadeiro) // not -> ! ou !=
fmt.Println(!false)
// FIM DOS LOGICOS

// 5 - UNARIOS
fmt.Println("-----------")
fmt.Println( "UNARIOS: ")
numeroUN := 10 // utilizado para aumentar, diminuir...
fmt.Println(numeroUN)
numeroUN ++ // numero = numero + 1
fmt.Println(numeroUN)
numeroUN -- // numero = numero - 1
fmt.Println(numeroUN)

numeroUN += 10 // numero = numero + 15
fmt.Println(numeroUN)
numeroUN -= 10 // numero = numero - 15
fmt.Println(numeroUN)

numeroUN *= 2 // numero = numero * 2
fmt.Println(numeroUN)
numeroUN /= 2 // numero = numero / 2
fmt.Println(numeroUN)
numeroUN %= 2 // numero = numero % 2
fmt.Println(numeroUN)

// FIM DOS UNARIOS

// 6 - TERNARIOS
fmt.Println("-----------")
fmt.Println( "TERNARIOS: ")
	numeroTER1 := 3
	numeroTER2 := 6
	var textoTER string

	if numeroTER1 > 5 {
		textoTER = "maior que 5"
	}else {
		textoTER = "menor que 5"
	}
	fmt.Println(textoTER)

	if numeroTER2 > 5 {
		textoTER = "maior que 5"
	}else {
		textoTER = "menor que 5"
	}
	fmt.Println(textoTER)
// FIM DOS TERNARIOS

}