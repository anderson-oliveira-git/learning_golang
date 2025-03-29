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

### 112. Estruturas de Controle

Go tem `if`, `for` e `switch`.

```go
if idade > 18 {
    fmt.Println("Maior de idade")
}

for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Este repositório será atualizado com mais conteúdo conforme avanço nos estudos sobre Go!
