# Array / Slice / Map

Neste módulo iremos analisar as formas de armazenar um conjunto de dados na 
linguagem Go.

### Array

Um array em Go é homogêneo e estático, ou seja, do mesmo tipo e de um tamanho fixo.
Vejamos abaixo como criar nosso primeiro array em Go.

```go
var notas [3]float64
notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
fmt.Println(notas)

// Ou na forma reduzida

notas := [3]float64{7.8, 4.3, 9.1}
fmt.Println(notas)
```

Você também pode declarar um array deixando o compilador de Go contar o 
tamanho do mesmo da seguinte forma:

```go
numeros := [...]int{1, 2, 3, 4, 5}
```

#### Percorrendo um array com for

Para percorrer uma array em Go podemos utilizar o for juntamente com uma palavra
reservada chamada "**range**", essa retorna o índice e o valor atual. Segue o exemplo: 

```go
numeros := [5]int{1, 2, 3, 4, 5}

for i, numero := range numeros {
    fmt.Printf("Indice: %d Valor: %d\n", i, numero)
}

// OU

for _, numero := range numeros {
    fmt.Printf("Valor: %d\n", numero)
}
```

### Slice

Diferente dos arrays, **o slice tem um tamanho dinâmico** e aponta por debaixo dos panos
para um array interno.

Logo, **slice não é um array**, mas sim define um pedaço de um array.

#### Declaração entre Array vs Slice

```go
a1 := [3]int{1, 2, 3} //array
s1 := []int{1, 2, 3}  //slice
```

#### Gerando um slice a partir de um array

No exemplo abaixo estou apontando para os valores do array do índice 1 até 2 (O último não é incluído).

```go
a2 := [5]int{1, 2, 3, 4, 5}
s2 := a2[1:3]
fmt.Println(a2, s2)
// [1 2 3 4 5] [2 3]
```

#### Gerando um slice a partir do método make

No exemplo abaixo, criei um slice de 10 elementos.

```go
s := make([]int, 10)
fmt.Println(s)
// [0 0 0 0 0 0 0 0 0 0]
```

Podemos também criar um slice de 10 elementos e com a capacidade do
array interno de 20.

```go
s := make([]int, 10, 20)
fmt.Println(s, len(s), cap(s)) // len - Tamanho do slice, cap - Capacidade do slice.
// [0 0 0 0 0 0 0 0 0 0] 10 20
```

#### Gerando slices a partir do append

```go
s1 := make([]int, 10, 20)
s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
fmt.Println(s1)
// [0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0]

s1 = append(s, 1) // Diferente do array, no slice não ocorre um erro quando se adiciona um valor extrapolando a capacidade máxima.
fmt.Println(s, len(s), cap(s)) // O slice internamente dobrou de tamanho para suportar mais um valor
//[0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0 1] 21 40
```

### Map

Map são uma estrutura de dados que armazenam valores no formato "chave:valor".

#### Criando um map em Go e adicionando valores

```go
aprovados := make(map[int]string)

aprovados[12345678] = "Maria"
aprovados[12345679] = "Ana"
aprovados[1234567213] = "Joao"

fmt.Println(aprovados)
//map[12345678:Maria 12345679:Ana 1234567213:Joao]
```

#### Deletando valores de um map

```go
delete(aprovados, 12345678)
fmt.Println(aprovados)
// map[12345679:Ana 1234567213:Joao]
```

#### Percorrendo um map em Go

Utilizando o map criado no exemplo anterior podemos percorrer ele usando o for e range.

```go
for cpf, nome := range aprovados {
    fmt.Printf("CPF: %d Nome: %s\n", cpf, nome)
}
```