package main

import "fmt"

func main()  {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	/*
		Ao mudar o valor inicial de um slice, o que está sendo modificado é o array interno.
		Logo, se eu mudar o valor da posição 0, ambos slices sofreram essa alteração.
	*/
	s1[0] = 7
	fmt.Println(s1, s2)
}
