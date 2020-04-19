package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	// Interface pode ser usado como um tipo genérico (Algo como o Object do Java)
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface {}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a linguagem do Google"}
	fmt.Println(coisa2)
}
