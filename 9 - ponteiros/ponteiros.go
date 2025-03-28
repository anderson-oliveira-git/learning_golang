package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 2
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória ::
	var variavel3 int
	var ponteiro *int // Declarando um ponteiro ::

	variavel3 = 100
	ponteiro = &variavel3

	// Caso queira ver o valor que está no ponteiro ::
	// será necessário usar a desreferenciação ::
	// exemplo: fmt.Println(*ponteiro) printa o ponteiro com um * ::
	fmt.Println(variavel3, *ponteiro)
}
