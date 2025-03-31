package main

import "fmt"

func main() {

	// Função anônima e autoexecutável ::
	retorno := func(messagem string) string {
		return fmt.Sprintf("Mensagem: %s", messagem)
	}("Mensagem passada como argumento para uma função anônima")

	fmt.Println(retorno)
}
