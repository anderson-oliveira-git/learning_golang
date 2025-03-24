package main

import "fmt"

func main() {
	// aritméticos: +, -, *, / e % ::
	soma := 1 + 2
	subtracao := 5 - 3
	divisao := 10 / 2
	multiplicacao := 20 * 3
	restoDivisao := 20 % 3

	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDivisao)

	// Não é possível fazer operações com variáveis de tipos diferentes ::
	// exemplo: var n1 int16 = 20 e var n2 int32 = 12 não podem sofrer operações ::

	// relacionais: >, >=, <, <=, == e != ::
	fmt.Println(2 > 1)
	fmt.Println(2 >= 2)
	fmt.Println(2 < 1)
	fmt.Println(2 <= 1)
	fmt.Println(2 == 1)
	fmt.Println(2 != 1)

	// lógicos: && (and) || (or) e ! (not)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false && true)

	// unários: ++, --, +=, -=, *=, /= e %= ::
	valor := 20

	valor++
	fmt.Println(valor)

	valor--
	fmt.Println(valor)

	valor += 3
	fmt.Println(valor)

	valor -= 2
	fmt.Println(valor)

	valor *= 5
	fmt.Println(valor)

	valor /= 2
	fmt.Println(valor)

	valor %= 2
	fmt.Println(valor)
}
