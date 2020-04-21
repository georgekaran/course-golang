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
    
