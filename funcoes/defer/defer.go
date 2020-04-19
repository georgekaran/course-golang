package main

import (
	"fmt"
)

func obterValorAprovado(n int) int {
	defer fmt.Println("Fim!")
	/*
	Atrasa a execução do código até o último momento de vida da função (Algo como o finally sem try catch)
	Esse tipo de recurso é interessante para fechar conexões com banco de dados
	*/
	if n > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return n
}

func main()  {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(300))
}
