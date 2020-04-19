package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //compilador conta o tamanho do array

	for i, numero := range numeros {
		fmt.Printf("Indice: %d Valor: %d\n", i, numero)
	}
	// OU
	for _, numero := range numeros {
		fmt.Printf("Valor: %d\n", numero)
	}
}
