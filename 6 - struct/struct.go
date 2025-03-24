package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	var e endereco
	// usuario ::
	u.nome = "Anderson"
	u.idade = 44

	// endereço ::
	e.logradouro = "Rua dos alfeneiros"
	e.numero = 123

	fmt.Println(u)

	// Usando um struct por inferência ::
	usuarioDois := usuario{"Davi", 21, e}

	fmt.Println(usuarioDois)
}
