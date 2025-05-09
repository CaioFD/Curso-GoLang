package main

import (
	"fmt"
	auxiliar "modulos/Auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main...")
	auxiliar.Escrever() //na main eu so consigo referenciar funcoes de outros pacotes como o auxiliar se ele comecar com letra Maiuscula.
	// mas no msm pacote eu consigo referenciar uma funcao comecando com letra minuscula, assim como esta no auxiliar.go
}