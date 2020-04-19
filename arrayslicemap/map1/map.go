package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// Mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[12345679] = "Ana"
	aprovados[1234567213] = "Joao"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("CPF: %d Nome: %s\n", cpf, nome)
	}

	fmt.Println(aprovados[12345678])
	delete(aprovados, 12345678)
	fmt.Println(aprovados[12345678])
}
