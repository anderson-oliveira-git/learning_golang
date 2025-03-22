package main

import (
	"fmt"
	"modulo/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	assistant.Escrever()

	cheking_email := checkmail.ValidateFormat("andersonreinilson@hotmail.com")

	// nil é a forma que go representa nulo que seria null em outras linguagens ::
	if cheking_email == nil {
		fmt.Println("O email informado é valido")
	} else {
		fmt.Println("O email informado é inválido")
	}
}
