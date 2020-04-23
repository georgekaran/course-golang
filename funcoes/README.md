# Funções

Funções em Go são pequenos blocos de código que tem como objetivo executar ações.

Entre as principais características de funções em Go podemos citar:
1. Podem retornar diversos valores ou não ter retorno.
2. Todas as variáveis da função devem ser declaradas.

### Criando uma função

#### Sem retorno

```go
// Sem retorno
func minhaFuncao(x int) {
    fmt.Println("Número:", x)
}
```

#### Com retorno

```go
// Simples
func minhaFuncao(x int) int {
    return x * 2
}

// Retorno composto
func minhaFuncao(x int) (int, string, bool) {
    return x * 2, "Estou multiplicando", true
}
```
