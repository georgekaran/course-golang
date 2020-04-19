package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo string
	turboLigado bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main()  {
	// Quando trabalho diretamente com o tipo, não necessito ter um tratamento especial para ligar o turbo.
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// Quando trabalho diretamente com a interface, neste caso, é necessário passar uma referência do ponteiro,
	// pois há método de receiver que fazer uso da referência da struct.
	// EVITAR ESSE TIPO DE CASO
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
