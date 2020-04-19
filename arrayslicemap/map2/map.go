package main

import "fmt"

func main()  {
	funcsESalarios := map[string]float64 {
		"José João": 11325.45,
		"Mauro Silva": 5153.13,
		"Alberto": 552.42,
	}

	funcsESalarios["Rafael Filho"] = 1350.0

	delete(funcsESalarios, "Não existe")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
