package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	primeiro = p1
	segundo = p2
	return // Retorno limpo, pois já declarei as variáveis de retorno
}

func main()  {
	segundo, primeiro := trocar(10, 5)
	fmt.Println(segundo, primeiro)
}
