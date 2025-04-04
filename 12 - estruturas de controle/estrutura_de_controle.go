package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	idade := 10

	/*
		Se estiver fazendo apenas uma única comparação,
		os () serão necessários, em Go, mas se estiver
		utilizando operadores lógicos, por exemplo,
		ai os () serão necessários ::
	*/
	if idade >= 18 {
		fmt.Printf("%d é considerado maior de idade", idade)
	} else {
		fmt.Printf("%d é considerado menor de idade", idade)
	}

	/*
		exemplo utilizando if init ::
		idadeInit só funciona no escopo do if em que ela foi inicializada ::
	*/
	if idadeInit := idade; idadeInit >= 18 {
		fmt.Printf("%d é considerado maior de idade", idadeInit)
	} else {
		fmt.Printf("%d é considerado menor de idade", idade)
	}

	// Switch
	diaDaSemana(2)
}

func diaDaSemana(dia int) string {
	switch dia {
	case 0:
		return "Domingo"
	case 1:
		return "Segunda-feira"
	case 2:
		return "Terça-feira"
	case 3:
		return "Quarta-feira"
	case 4:
		return "Quinta-feira"
	case 5:
		return "Sexta-feira"
	case 6:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}
