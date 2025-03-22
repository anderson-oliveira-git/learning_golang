package assistant

import "fmt"

// Quando a função é escrita com a primera letra maiuscula ::
// essa função se comporta como sendo publica ::
// sendo possível acessá-la em outros scripts ::
func Escrever() {
	fmt.Println("Escrevendo da função assistante")
	somar()
}

// Quando a função é escrita com a primera letra minuscula ::
// essa função se comporta como sendo privada ::
// sendo possível acessá-la apenas no escopo onde foi criada ::
func somar() {
	fmt.Println("Somando valores...")
}
