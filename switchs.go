package main

import "fmt"

func main() {

	switch modulo := 4 % 2; modulo {

	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("No es par")
	}

	//Switch sin condicion

	value := 0
	switch {
	case value == 200:
		fmt.Println("Es igual a 200")
	case value > 200:
		fmt.Println("Es mayor a 200")
	default:
		fmt.Println("Ninguan condicion se cumple")
	}
}