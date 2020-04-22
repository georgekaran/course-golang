# Fundamentos

Neste módulo iremos analisar os fundamentos da linguagem Go, 
como formas de declarar uma variável, tipos, funções, etc...

### Primeiro programa em Go

Para criarmos o nosso primeiro programa em Go, vou iniciar pela criação um arquivo 
chamado "primeiro.go".

Neste arquivo, escrevi o seguinte código abaixo:
```go
package main

import "fmt"

func main() {
	fmt.Println("Primeiro Programa")
}
```

Explicações:
1. **package main** - Programas executáveis iniciam pelo pacote main (No capítulo 
sobre pacotes vamos trabalhar mais sobre a manipulação destes)
2. **import "fmt"** - Os códigos em GO são organizados em pacotes e para usá-los 
é necessário declarar um ou vários imports.
3. **func main()** - A função main é a porta de entrada de um programa GO 
(Similar ao public static void main do Java).

### Comandos

Antes de nos aprofundarmos na linguagem, há alguns comandos úteis do Go no terminal que é importante
conhecer.

1. Executar um arquivo .go:
    ```shell script
    go run ARQUIVO.go
    ```

2. Examina um arquivo .go e verifica se há códigos suspeitos 
(falhas, falta de argumentos):
    ```shell script 
    go vet ARQUIVO.go
    ```

3. Gerar um executável a partir do .go:
    ```shell script
    go build ARQUIVO.go
    ```

4. Instala pacotes na sua máquina (Dentro do GOPATH) 
(-u é para atualizar versões já existentes se necessário):
    ```shell script
    go get -u ENDERECO_DO_PACOTE
    Ex:
    go get -u "github.com/go-sql-driver/mysql"
    ```
    
### Constantes e Variáveis

A linguagem Go é fortemente tipada e por esse motivo todas as variáveis e constantes
devem possuir um tipo, seja ele explícito ou implícito. 

Declarando uma constante informando o tipo:

```go
const PI float64 = 3.1415
```

Declarando uma constante sem informar o tipo (Neste caso, o tipo será inferido pelo
compilador):

```go
const PI = 3.1415 // float64
```

O mesmo exemplo pode ser através da variável:

```go
var raio = 3.1 // float64
```

### Const vs Var

Seguindo a mesma linha do Javascript, no Go uma **constante não pode sofrer alterações
de valor** e o **valor deve ser definido na sua declaração** como foi feito nos exemplos
acima.

Já uma variável pode ser declarada sem receber valor e sofrer alterações durante 
o curso de sua aplicação. Exemplo:

```go
var raio float64 // default é 0
raio = 3.1
```

### Formas de declarar uma variável

Em Go há duas formas de se declarar variáveis, são elas:

```go
var raio float64 = 3.2

// Forma reduzida de criar uma var
raioReduzido := 3.2
```

É importante ressaltar que a forma reduzida somente **cria variáveis e 
não constantes!**

### Tipos em Go

**Inteiros positivos**: uint8, uint16, uint32, uint64. Exemplo:

```go
var inteiroPositivo uint32 = 1
fmt.Println("Valor:", inteiroPositivo, "Tipo:", reflect.TypeOf(inteiroPositivo))
// Valor: 1 Tipo: uint32
```

**Inteiros com sinal**: int, int8, int16, int32, int64: Exemplo:

Por padrão se você estiver trabalhando com inteiros use o "**int**".

```go
var inteiro = -90
fmt.Println("Valor:", inteiro, "Tipo:", reflect.TypeOf(inteiro))
// Valor: -90 Tipo: int
```

**Byte**: byte (é na verdade um alias para uint8). Exemplo:

```go
var bb byte = 254
fmt.Println("Valor:", bb, "Tipo:", reflect.TypeOf(bb))
// Valor: 254 Tipo: uint8
```

**Rune**: rune (representa um mapeamento da tabela Unicode (int32)). Exemplo:

```go
var r rune = 'a'
fmt.Println("Valor:", r, "Tipo:", reflect.TypeOf(r))
// Valor: 97 Tipo: int32
```

**Números reais**: float32, float64. Exemplo:

```go
var decimal float64 = 49.90
fmt.Println("Valor:", decimal, "Tipo:", reflect.TypeOf(decimal))
// Valor: 49.9 Tipo: float64
```

**Boolean**: bool. Exemplo:

```go
var b1 bool = true
fmt.Println("Valor:", b1, "Tipo:", reflect.TypeOf(b1))
// Valor: true Tipo: bool
```

**String**: string. Exemplo:

```go
s1 := "Aprendendo GO!"
fmt.Println("Valor:", s1, "Tipo:", reflect.TypeOf(s1))
// Valor: Aprendendo GO! Tipo: string

// String com múltiplas linhas
s2 := `Estou
escrevendo
em múltiplas
linhas`
fmt.Println("Valor:", s2, "Tipo:", reflect.TypeOf(s2))
// Valor: Estou
// escrevendo
// em múltiplas
// linhas Tipo: string
```

**Char**: int32. Exemplo:

```go
char := 'a'
fmt.Println("Valor:", char, "Tipo:", reflect.TypeOf(char))
// Valor: 97 Tipo: int32
```

### Valores defaults

Como vimos anteriormente, em Go quando se declara uma variável não é necessário
atribuir um valor para o mesmo contanto que o tipo seja informado. Mas afinal qual
é o valor default para cada tipo?

```go
var a int // 0
var b float64 // 0
var c bool // false
var d string // ""
var e *int // <nil>

fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e)
// 0 0 false "" <nil>
```

### Conversão entre tipos básicos

#### Convertendo valores para decimal:

```go
inteiro := 1
decimal := float64(inteiro)
fmt.Println("Valor:", decimal, "Tipo:", reflect.TypeOf(decimal))
// Valor: 1 Tipo: float64
```

#### Convertendo valores para inteiro:

```go
decimal := 1.5
inteiro := int(decimal) // Arredondamento é feito para baixo...
fmt.Println("Valor:", inteiro, "Tipo:", reflect.TypeOf(inteiro))
// Valor: 1 Tipo: int
```

#### Convertendo inteiro para string:

```go
import "strconv"
...
str := strconv.Itoa(123)
fmt.Println("Valor:", str, "Tipo:", reflect.TypeOf(str))
// Valor: 123 Tipo: string
```

#### Convertendo string para inteiro:

```go
import "strconv"
...
num, _ := strconv.Atoi("123") // Esse método retorna dois valores, o primeiro é o número convertido e o segundo é um erro caso o número seja inválido.
fmt.Println("Valor:", num, "Tipo:", reflect.TypeOf(num))
// Valor: 123 Tipo: int
```

#### Convertendo string para boolean:

```go
import "strconv"
...
b, _ := strconv.ParseBool("true") // Esse método retorna dois valores, o primeiro é o número convertido e o segundo é um erro caso o número seja inválido.
fmt.Println("Valor:", b, "Tipo:", reflect.TypeOf(b))
// Valor: true Tipo: bool
```

### Criando uma função básica

Sabemos que para o ponto de partida de um programa Go é a função main, mas podemos
criar outras funções para fazer pequenas atividades em blocos.

```go
package main

import "fmt"

// Devemos explicitamente declarar os parâmetros e o retorno de uma função
func somar(x int, y int) int {
	return x + y
}

func main() {
    fmt.Println("Soma:", somar(5, 10))
}
```

Este foi um pequeno exemplo de funções, ainda teremos um capítulo completo sobre
esse assunto.

### Ponteiros em Go

OBS: Diferente de outras linguagens, em Go não é possível fazer aritmética em cima de 
ponteiros.

Declarando um ponteiro em Go:

```go
var p *int = nil
```

Pegando a referência de memória e alterando o valor através de ponteiro:

```go
i := 1
var p *int = nil

p = &i // Aqui estou pegando o endereço de memória da variável i.
fmt.Println(p, *p, i, &i)
// 0xc0000a4010 1 1 0xc0000a4010

*p++ // Desreferenciando (pegando o valor)
fmt.Println(p, *p, i, &i)
// 0xc0000a4010 2 2 0xc0000a4010
```