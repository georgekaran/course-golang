# Estruturas de controle

Neste módulo iremos analisar as estruturas de controle da linguagem Go, 
como if, else, for, switch, etc...

### IF

A estrutura de controle "if" por padrão **não possui parênteses** 
(salvo casos com condições aninhadas) e mesmo que o bloco
seja de 1 linha associada **é necessário utilizar as chaves**.

#### IF ELSE

```go
nota := 7.2
if nota >= 7 {
    fmt.Println("Aprovado!!!")
} else {
    fmt.Println("Reprovado!!!")
}
```

#### IF ELSE IF

Aqui temos o exemplo acima utilizando if, else if, ...

```go
nota := 7.2
if nota >= 9 {
    fmt.Println("Passou muito bem!!!")
} else if nota >= 8 && nota < 9  {
    fmt.Println("Passou bem!!!")
} else if nota >= 7 && nota < 8  {
    fmt.Println("Passou!!!")
} else {
    fmt.Println("Reprovou!!!")
}
```

#### IF com init

No Go é possível iniciar uma variável a partir de uma condição if, essa variável
estará disponível apenas ao escopo do bloco em que foi iniciado. Vamos com um exemplo:

```go
func numeroAleatorio() int {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)
    return r.Intn(10)
}

func main() {
    //...
    if i := numeroAleatorio(); i > 5 { // também é suportado no switch 
        fmt.Println("Ganhou!!!")
    } else {
        fmt.Println("Perdeu!!!")
    }
}
```

### FOR

Diferente de outras linguagens, **no Go a única estrutura de repetição disponível é
o for**. Apesar de não existir o while conseguimos atingir os mesmos resultados
utilizando o "for" no Go.

#### FOR - While

```go
i := 1
for i <= 10 {
    fmt.Println(i)
    i++
}
```

#### FOR - Default

```go
for i := 0; i <= 20; i += 2 {
    fmt.Printf("%d\n", i)
}
```

#### FOR - Infinito

```go
for {
    fmt.Println("Para sempre...")
    time.Sleep(time.Second) // Timeout de 1 segundo
}
```

### SWITCH

A estrutura de controle "switch" em Go é similar a padrões de outras linguagens já
conhecidas como Javascript, Java, etc... **A diferença se da por não existir a palavra
"break"**, ou seja, quando um case do switch é verdadeiro automaticamente o Go já sai
do bloco do switch.

#### SWITCH - Com base em uma variável

```go
nota := 9
switch nota {
case 10:
    /*
        Neste caso como estou utilizando a palavra reservada fallthrough
        Mesmo que a nota fosse 10 o switch continuaria executando até achar outra
        condição que seja verdadeira.
    */
    fallthrough
case 9:
    fmt.Println("A")
case 8:
    fmt.Println("B")
case 6, 7: // É possível utilizar mais de uma condição para um case
    fmt.Println("C")
default:
    fmt.Println("Nota inválida")
}
```

#### SWITCH - Sem variável (True)

```go
t := time.Now()
switch { // switch true 
case t.Hour() < 12:
    fmt.Println("Bom dia")
case t.Hour() < 18:
    fmt.Println("Boa tarde")
default:
    fmt.Println("Boa noite")
}
```