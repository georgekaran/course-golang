package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, ch chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				ch <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	// É importante fechar o canal para não gerar nenhum deadlock
	close(ch) // Estou dizendo que não vai recebido mais nenhum dado no canal
}

func main() {
	c := make(chan int, 60)
	go primos(cap(c), c)
	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim")
}
