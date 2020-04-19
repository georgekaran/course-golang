package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s)) //cap: Capacidade máxima do slice

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // Diferente do array, no slice não ocorre um erro quando se adiciona um valor extrapolando a capacidade máxima.
	fmt.Println(s, len(s), cap(s)) // O slice internamente dobrou de tamanho para suportar mais um valor
}
