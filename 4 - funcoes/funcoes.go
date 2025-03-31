package main

import "fmt"

func somarDoisNumeros(numeroUm int8, numeroDois int8) int8 {
	return numeroUm + numeroDois
}

// caso os parâmetros sejam do mesmo tipo, basta informar apenas para o último ::
// as funcões em Go podem ter mais de um retorno ::
func calculoMatematico(numeroUm, numeroDois int8) (int8, int8) {
	soma := numeroUm + numeroDois
	subtracao := numeroUm - numeroDois

	return soma, subtracao
}

func main() {
	var total int8 = somarDoisNumeros(10, 20)
	fmt.Println(total)

	var fn = func(messagem string) {
		fmt.Println(messagem)
	}

	fn("Mensagem passada por argumento")

	// caso não queira utilizar um dos retornos basta passar um underline ::
	// exemplo: resultadoCalculoSoma, _ := calculoMatematico(10, 15) ::
	// no exemplo acima, apenas a soma será retornada ::
	resultadoCalculoSoma, resultadoCalculoSubtracao := calculoMatematico(10, 15)
	_, subtracao := calculoMatematico(10, 15)

	fmt.Println(resultadoCalculoSoma, resultadoCalculoSubtracao)
	fmt.Println(resultadoCalculoSoma, subtracao)
}
