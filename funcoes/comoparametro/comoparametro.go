package main

import "fmt"

func multiplicao(x, y int) int  {
	return x * y
}

func exec(funcao func(int, int) int, p1, p2 int) int  {
	return funcao(p1, p2)
}

func main() {
	fmt.Println("Resultado", exec(multiplicao, 10, 5))
}