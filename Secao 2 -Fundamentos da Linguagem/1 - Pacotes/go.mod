module modulos

go 1.24.3

// importando pacotes externos 
//go.mod → diz quais pacotes você usa e em qual versão.
//go.sum → garante que essas versões não mudaram ou foram corrompidas.

require github.com/badoux/checkmail v1.2.4 // indirect
// caso queira remover os pacotes externsos, usar o comando go mod tydi
