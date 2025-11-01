// comando para gerar um modulo: go mod init aula-3
// comando para incluir dependência: go get <nome-da-dependencia> <github.com/badoux/checkmail>

package main

import (
	"fmt"
	checkmail "github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Adicionando módulo em Go!")
	result := checkmail.ValidateFormat("teste@teste.com")
	fmt.Println("Resultado da validação do email:", result)
}
