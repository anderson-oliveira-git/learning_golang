package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos inteiros ::
	// int, int8, int16, int32, int64 ::
	// uint - aceita apenas valores positivos e segue a mesma lógica do int com 8, 16, 32 ou 64 ::
	// rune - é um apelido caso queira usar int32 ::
	// byte - é um apelido caso queira usar int8 ::
	// O int sem o valor a direita, utiliza a base do do processador como valor ::
	// exemplo: Se o pc for base x64 esse int será int64 ::

	var numeroUm int64 = -1000000000000000000
	fmt.Println(numeroUm)

	// aceita apenas números positivos ::
	var numeroDois uint32 = 1000000000
	fmt.Println(numeroDois)

	// rune é um alias para o int32 ::
	var numeroTres rune = 1000000000
	fmt.Println(numeroTres)

	// byte é um alias para o int8 ::
	var numeroQuatro byte = 100
	fmt.Println(numeroQuatro)

	// tipo float ::
	// são apenas dois tipos: float32 e float64 ::

	var numeroRealUm = 123.34354534534534534
	fmt.Println(numeroRealUm)

	var numeroRealDois = 123.34354534534534534
	fmt.Println(numeroRealDois)

	// tipo string ::

	// em Go sempre usa aspas duplas para strings ::
	var nome string = "Anderson"
	fmt.Println(nome)

	// o tipo 'char' não existe em Go, mas usando aspas simples ele retorna o número desse caractere na tabela ASCII ::
	sexo := 'M'
	fmt.Println(sexo)

	// M na tabela ASCII é igual a 77 ::
	if sexo == 77 {
		fmt.Println("Essa letra é a M")
	} else {
		fmt.Println("Essa letra não é a M")
	}

	// obs: todo tipo de dado tem seu valor '0' ::
	// tipo int e float recebem 0 caso não atribua nada ::
	// tipo string recebe string vazia ::

	var maiorDeIdade bool = true
	fmt.Println(maiorDeIdade)q

	// o tipo error precisa da pacote errors ::
	var erro error = errors.New("Erro interno!")
	fmt.Println(erro)
}
