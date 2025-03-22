package main

import "fmt"

func main() {
	// Forma de criar variáveis atribuindo o tipo de forma explicita ::
	var nome string = "Anderson"

	// Forma de criar variáveis atribuindo o tipo por inferência ::
	idade := 32

	// Forma de criar várias variáveis de uma vez ::
	// O valor pode ser atribuido dentro do var () ou fora ::
	// Exemplo
	var (
		email string
		cpf   string
		rg    string
	)

	// Forma de criar várias variáveis por inferência ::
	telefone, celular := "(00) 0 0000-0000", "(00) 0 0000-0000"

	// Forma como a linguagem GO inverte o valor entre duas variáveis ::
	telefone, celular = celular, telefone

	email = "andersonreinilson@hotmail.com"
	cpf = "111.111.111-11"
	rg = "233.233.233"

	// Forma de declarar constantes ::
	const PI = 3.1415

	fmt.Println(nome, idade)
	fmt.Println(email, cpf, rg, telefone, celular)
}
