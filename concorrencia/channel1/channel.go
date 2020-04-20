package main

import "fmt"

func main() {
	// Criando um canal de comunicação que enviará dados inteiros.
	ch := make(chan int, 1)

	// Enviando dados para o canal (escrita)
	ch <- 1
	// Recebendo dados para o canal (leitura)
	<- ch

	ch <- 2
	fmt.Println(<-ch)
}
