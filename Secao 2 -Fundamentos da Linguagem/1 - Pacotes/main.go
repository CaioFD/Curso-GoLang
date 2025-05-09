package main

import (
	"fmt"
	auxiliar "modulos/Auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main...")
	auxiliar.Escrever() //na main eu so consigo referenciar funcoes de outros pacotes como o auxiliar se ele comecar com letra Maiuscula.
	// mas no msm pacote eu consigo referenciar uma funcao comecando com letra minuscula, assim como esta no auxiliar.go

	erro := checkmail.ValidateFormat("cursodeGO@gmail.com") // Pacote Externo de checkmail, usado a partir do go get github.com/badoux/checkmail@latest e adicionado automaticamente no go.mod e go.sum
	fmt.Println(erro)
	// caso retorne "nulo" (<nil>), significa que o email e valido
	// caso o contráio irá retornar "invalid format"
}