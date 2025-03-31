package main

import (
	"fmt"
	"time"
)

func main() {

	// em Go não existe while nem do while, existe apenas for ::
	for j := 0; j < 10; j++ {
		fmt.Printf("Incrementando i: %d \n", j)
		time.Sleep(time.Second)
	}

	// for com range ::
	nomes := [3]string{
		"Anderson",
		"Silva",
		"Oliveira",
	}

	for _, nome := range nomes {
		fmt.Printf("Nome: %s \n", nome)
	}
}
