package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     byte
	altura    uint16
}

type estudante struct {
	// chamar uma struct dentro da outra, simula a herança entre classes ::
	// isso tras todos os campos da struct pessoa, para dentro de estudante ::
	dadosPessoais pessoa
	curso         string
	faculdade     string
}

func main() {
	p1 := pessoa{"Anderson", "Oliveira", 35, 176}
	e1 := estudante{p1, "Engenharia de Software", "Estácio"}
	fmt.Println(p1, e1)
}
