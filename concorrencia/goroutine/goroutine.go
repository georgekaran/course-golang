package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qnt int) {
	for i := 0; i < qnt; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i + 1)
	}
}

func main() {
	//fale("Maria", "Por que você não fala comigo?", 3)
	//fale("João", "So posso falar depois de você!", 1)

	// Ao usar a palavra reservada "go" a instrução será executada de forma independente
	// Também ao usar a palavra go é criado uma go routine.
	// Algo similar a uma Thread, mas muito mais leve.
	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)
	// Atenção: Uma go routine só executada se a thread principal ainda estiver em execução.

	go fale("Maria", "Ei...", 10)
	fale("João", "Opa...", 5)
}
