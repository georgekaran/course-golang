# Testes

### Criando um teste em GO

Para criar o primeiro teste em GO vamos começar criando uma função que calcula
a média de entre notas.

```go
package matematica

func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
```

OBS: É importante que a função de calcular a média seja iniciada com letra maiúscula
assim, conseguimos ter acesso fora do pacote.

Em seguida, para que o nosso teste seja considerado pelo Go é necessário 
seguir as regras abaixo:

1. O nome do arquivo deve conter no final "_test.go".
2. A função teste dentro do arquivo deve começar por "Test" e receber como 
parâmetro um ponteiro para "testing.T". Para o nosso caso, seria algo nesse sentido:
    ```go
   func TestMedia(t *testing.T) {...}
    ```

Com base nas regras acima, podemos criar nosso primeiro teste em Go.

```go
package matematica

import "testing"

const errorPadrão = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(errorPadrão, valorEsperado, valor)
	}
}
```

### Executando

Para executar os testes, você deve entrar na pasta onde está localizado seu arquivo
e executar o comando abaixo:
```shell script
go test
```

Ou se preferir executar um conjunto de testes:
```shell script
go test ./...
```

### Observações

* Caso você tenha uma suíte de testes grande, talvez seja necessário rodar os testes de
forma paralela. Para fazer isso, basta adicionar a seguinte linha de código no ínicio
de seu teste:
```go
func TestAbc(t *testing.T) {
    t.Parallel()
    ...
}
```

* Para pular um teste:
```go
    t.Skip("Pulei este teste devido...")
```

* Para forçar um erro no teste:
```go
    t.Fail()
```

* Verificar a cobertura dos testes:
```shell script
    go test -cover
```

* Verificar a cobertura dos testes:
```shell script
    go test -cover
```

* Gerando cover profile e exibindo em um .html:
```shell script
    go test -coverprofile=resultado.out
    go tool cover -html=resultado.out
```