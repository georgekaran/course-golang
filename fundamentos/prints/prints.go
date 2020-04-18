package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println(" linha.")

	x := 3.1415
	fmt.Println("O valor de x é", x)

	xs := fmt.Sprint(x) // Converte para String
	fmt.Println("O valor de x é " + xs)

	fmt.Printf("O valor de x é %.2f.", x) // Print com ganchos e formato para apenas 2 casas decimais

	a := 1
	b := 1.999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
