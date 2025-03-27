package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	// primeira forma de declarar um array ::
	// nome_variável [quantidade de items do array] tipo ::
	// se o tamanhno do array não for especificado, ele vira um slice ::
	var nomes [5]string

	// atribuindo valores para cada indice do array ::
	nomes[0] = "Anderson"
	nomes[1] = "Oliveira"
	nomes[2] = "Go"
	nomes[3] = "PHP"
	nomes[4] = "Typescript"

	fmt.Println(nomes)

	// segunda forma de declarar um array ::
	frutas := [3]string{"banana", "abacate", "pera"}

	fmt.Println(frutas)

	// trabalhando com slices ::

	// essa estrutura possui tamanho dinâmico, mas ainda é limitado ao seu próprio tipo ::
	slice := []int{10, 11, 12, 13, 14, 15}
	slice = append(slice, 16)

	frutas2 := frutas[0:2]

	fmt.Println(frutas2)
}
