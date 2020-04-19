package main

import "fmt"

// Posso ter uma função init para cada arquivo
func init()  {
	fmt.Println("Inicializando...")
}

func main()  {
	fmt.Println("Main...")
}