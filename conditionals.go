package main

import "fmt"

func main() {
	//Condicionales If

	valor1 := 1
	valor2 := 3

	if valor2 == 3 && valor1 == 1 {
		fmt.Println("El valor es 3 y el numero es 1")
	} else {
		fmt.Println("El valor no es 3")
	}

	//Condicionales if con un OR

	if valor1 == 3 || valor2 == 2 {
		fmt.Println("Se cumple una condicion ")
	} else {
		fmt.Println("No se cumple alguna condicion")
	}

}
