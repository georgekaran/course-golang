package main

import "fmt"

func main() {
	fmt.Printf("Outro programa em %s!\n", "GO")
}

/*
	go vet ARQUIVO: Examina um arquivo .go e verifica se há códigos suspeitos (falhas, falta de argumentos).
	go run ARQUIVO: Executa um programa .go.
	go build ARQUIVO: Gera um executável apartir do .go.
	go get -u ENDERECO_DO_PACOTE: Instala pacotes Go na sua máquina (Dentro do GOPATH).
*/