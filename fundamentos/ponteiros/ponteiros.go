package main

import "fmt"

func main() {
	i := 1

	fmt.Println(i)

	var p *int = nil
	p = &i // Pegar referência de memória da variável i
	*p++ // Desreferenciando (pegando o valor)
	i++

	// Go não tem aritmética de ponteiros
	// p++ - Não pode

	fmt.Println(p, *p, i, &i)
}
