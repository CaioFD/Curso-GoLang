package main

import "fmt"

func somar(numero1 int8, numero2 int8) int8 { // Caso a funcao tenha um retorno, voce especifica esse retorno depois do parentese
	return numero1 + numero2
}

// No GO as funcoes podem ter mais de um retorno

//podemos declara o tipo do parametro dessa forma, declarando apenas o ultimo (Caso os dois sejam do mesmo tipo)
func calculos(n1, n2 int8) (int8, int8) { // Na funcao calculos ela recebe dois parametros do tipo int8(n1 e n2)+ dois retornos ao inves de um so
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da funcao 1")
	fmt.Println(resultado)

	fmt.Println("\n")

	resultadoSoma, resultadoSubtracao := calculos(10, 8)
	fmt.Println("Resultado da Soma: ", resultadoSoma, "\n","Resultado da Subtracao: ", resultadoSubtracao)

	fmt.Println("\n")

	// Caso vc nao queira ter o resultado de ou da soma ou da subtracao basta colocar um "_" no lugar de uma delas
	resultadoSoma2, _ := calculos(10, 8)
	fmt.Println("Resultado da Soma: ", resultadoSoma2)

	_, resultadoSubtracao2 := calculos(10, 8)
	fmt.Println("Resultado da Subtracao: ", resultadoSubtracao2)

	
}