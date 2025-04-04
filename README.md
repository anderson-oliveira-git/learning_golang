# Estudos sobre Go

Este repositório contém anotações e exemplos de código sobre conceitos fundamentais da linguagem Go.

## Tópicos

### 1. Packages

Em Go, o código é organizado em pacotes (packages), que permitem modularização e reutilização de código. O ponto de entrada de um programa é o pacote `main`.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 2. Variáveis

Variáveis podem ser declaradas de diferentes formas em Go:

```go
var nome string = "Alice"
idade := 25 // Inferência de tipo
const PI = 3.14
```

### 3. Tipos de Dados

Go possui tipos pré-definidos como `int`, `float64`, `string`, `bool`, entre outros.

```go
var numero int = 10
var decimal float64 = 3.14
var ativo bool = true
```

### 4. Funções

As funções são definidas usando `func` e podem retornar valores.

```go
func soma(a int, b int) int {
    return a + b
}
```

### 5. Operadores

Go suporta operadores aritméticos, relacionais e lógicos.

```go
soma := 10 + 5
comparacao := 10 > 5
logico := true && false
```

### 6. Struct

Structs permitem criar tipos personalizados.

```go
type Pessoa struct {
    Nome  string
    Idade int
}
```

### 7. Simulando Herança

Go não suporta herança, mas é possível compor structs.

```go
type Endereco struct {
    Rua string
}

type Pessoa struct {
    Nome     string
    Endereco // Composição
}
```

### 8. Arrays e Slices

Arrays possuem tamanho fixo, enquanto slices são dinâmicos.

```go
var arr [3]int = [3]int{1, 2, 3}
slice := []int{1, 2, 3}
```

### 9. Ponteiros

Ponteiros permitem manipular endereços de memória.

```go
var x int = 10
p := &x // Ponteiro para x
```

### 10. Arrays Internos

Arrays internos podem ser usados dentro de structs para armazenar múltiplos valores fixos.

```go
type Matriz struct {
    Dados [3][3]int
}
```

### 11. Maps

Maps são coleções chave-valor.

```go
pessoa := map[string]int{"Alice": 25, "Bob": 30}
```

### 12. Estruturas de Controle

Go tem `if`, `for` e `switch`.

```go
if idade > 18 {
    fmt.Println("Maior de idade")
}

for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 13. Loops em Go

Em Go, o loop `for` é a única estrutura de repetição disponível, mas ele pode ser usado de diferentes formas:

#### Loop Clássico
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

#### Loop como While
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

#### Loop Infinito
```go
for {
    fmt.Println("Executando...")
    break // Necessário para evitar loop infinito
}
```

#### Iterando sobre um Slice
```go
numeros := []int{10, 20, 30, 40}
for i, num := range numeros {
    fmt.Printf("Índice: %d, Valor: %d\n", i, num)
}
```

#### Iterando sobre um Map
```go
idades := map[string]int{"Alice": 25, "Bob": 30}
for nome, idade := range idades {
    fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
}
```

### 14. Funções com Retorno Nomeado em Go

Em Go, é possível definir funções com retornos nomeados. Isso significa que os valores de retorno podem ser declarados junto com os tipos na assinatura da função, permitindo que sejam atribuídos dentro da função sem a necessidade de um `return` explícito.

#### Exemplo de Função com Retorno Nomeado

```go
package main

import "fmt"

// Função com retorno nomeado
func somaESubtracao(a, b int) (soma int, subtracao int) {
    soma = a + b
    subtracao = a - b
    return // Retorna os valores nomeados automaticamente
}

func main() {
    s, sub := somaESubtracao(10, 5)
    fmt.Printf("Soma: %d, Subtração: %d\n", s, sub)
}
```

### 15. Funções Variádicas em Go

Em Go, funções variádicas permitem que um número indefinido de argumentos do mesmo tipo seja passado para a função. Isso é útil quando não se sabe quantos argumentos serão fornecidos durante a chamada da função.

#### Exemplo de Função Variádica

```go
package main

import "fmt"

// Função variádica que soma todos os números passados como argumento
func soma(numeros ...int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}

func main() {
    resultado := soma(1, 2, 3, 4, 5)
    fmt.Printf("Soma: %d\n", resultado)
}
```

### 16. Funções Anônimas em Go

Funções anônimas são aquelas que não possuem um nome e podem ser usadas para criar funções "inline" ou para passar funções como parâmetros. Elas são muito úteis quando você precisa de uma função rapidamente e não quer se preocupar em dar um nome a ela.

#### Exemplo de Função Anônima

```go
package main

import "fmt"

func main() {
    // Função anônima atribuída a uma variável
    soma := func(a, b int) int {
        return a + b
    }

    // Chamando a função anônima
    resultado := soma(10, 20)
    fmt.Printf("Soma: %d\n", resultado)

    // Usando a função anônima diretamente dentro de outra função
    resultado2 := func(a, b int) int {
        return a * b
    }(5, 6)

    fmt.Printf("Multiplicação: %d\n", resultado2)
}


Este repositório será atualizado com mais conteúdo conforme avanço nos estudos sobre Go!
