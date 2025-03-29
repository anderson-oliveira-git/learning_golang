package main

import "fmt"

func main() {
	fmt.Println("Maps")

	/*
		Maps funcionam como objetos em Javascript ::
		o tipo que fica entre [] é o tipo da chave e ::
		o tipo que fica fora dos [] é o tipo do valor ::
		todas as chaves e valores precisam ser do mesmo tipo que foram definidas ::
		exemplo: map[int]string não pode ter a chave com "" e nem o valor sendo inteiro ::
	*/
	usuario := map[string]string{
		"nome":      "Anderson",
		"sobrenome": "Oliveira",
	}

	// Para acessar o valor do map, funciona exatamente da mesma forma como acessa um dado de um array ::
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	// também é possível aninhar maps ::
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro_nome": "Anderson",
			"ultimo_nome":   "Oliveira",
		},
	}

	fmt.Println(usuario2["nome"]["primeiro_nome"])

	// Usando a função delete(), é possível remover items do map ::
	delete(usuario2, "nome") // isso remove a chave nome do map usuario2 ::

	// caso queira adicionar, basta fazer igual nos arrays ::
	usuario2["nome"] = map[string]string{
		"primeiro_nome": "Maria",
		"ultimo_nome":   "Bairro",
	}

	fmt.Println(usuario2["nome"]["primeiro_nome"])
}
