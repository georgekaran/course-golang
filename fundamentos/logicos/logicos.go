package main

import "fmt"

func compras(trab1, trab2  bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSoverte := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSoverte
}

func main()  {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
