package main

import "fmt"

func main() {
	fmt.Println("Arrays internos")

	// A função make() recebe 3 argumentos ::
	// o tipo do slice, o tamanho do slice e a quantidade máxima ::
	slice := make([]float32, 10, 15)

	fmt.Println(slice)
	fmt.Println(len(slice)) // A função len() mostra o tamanho do slice ::
	fmt.Println(cap(slice)) // A função cap() mostra a capacidade total ::
}
