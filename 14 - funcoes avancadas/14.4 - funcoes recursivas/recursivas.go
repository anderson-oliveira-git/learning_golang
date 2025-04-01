package main

import "fmt"

func main() {
	fmt.Println("Funções recursivas")
	posicao := uint(10)

	fmt.Println(fibonacci(posicao))
}

/*
	A função recursiva necessita de uma condição de parada,
	assim evitando estouro de pilha. ::
*/
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
