package main

import "fmt"

func main() {
	total := somarNumeros(1, 2, 3, 4, 5)
	fmt.Println(total)
}

// usando ...{tipo} fará com que a função receba vários argumentos ::
// a função variátiva aceita apenas um único parâmetro variático e ele sempre precisa ser o último parâmetro ::
func somarNumeros(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
