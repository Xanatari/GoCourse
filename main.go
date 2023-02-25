package main

import (
	"fmt"
)

func main() {

	// Declaracion variables
	const num1 float64 = 3.14
	const num2 = 3.1416

	//Declaracion de Valriables enteras

	base := 12
	var altura int = 14
	var area int

	//Zero values

	//default value 0
	var a int

	//Default Value 0
	var b float64
	//Default vacio
	var c string
	//Default false
	var d bool

	fmt.Println("The number one is ", num1)
	fmt.Println("The number two is ", num2)
	fmt.Println("The base is ", base)
	fmt.Println("Altura", altura)
	fmt.Println("Area is", area)
	fmt.Println(a, b, c, d)

	var basecuadrado = 10
	quadrature := basecuadrado * basecuadrado
	fmt.Println("Cuadrado--")
	fmt.Println(quadrature)

	resta := basecuadrado - basecuadrado
	fmt.Println("Resta--")
	fmt.Println(resta)

	suma := basecuadrado + basecuadrado + 1
	fmt.Println("Suma--")
	fmt.Println(suma)

	divicion := basecuadrado / basecuadrado
	fmt.Println("Divicion--")
	fmt.Println(divicion)

	residiu := basecuadrado / basecuadrado
	fmt.Println(residiu)

	basecuadrado++
	fmt.Println("Incremental")
	fmt.Println(basecuadrado)

	fmt.Println("Decrecimiento")
	basecuadrado--
	fmt.Println(basecuadrado)

	fmt.Printf("Nosotros los %d de la ciudad de Bogota estamos en guerra con Cali aprox %d soldados ", basecuadrado, suma)

	fmt.Printf("basecuadrado: %T ", basecuadrado)

	fmt.Println("Condicional")
	if suma == 20 {
		fmt.Println(suma + divicion)
		d = true

	} else {
		fmt.Println(divicion)
	}

	if d {
		fmt.Println("True-->")
		fmt.Println(d)
	} else {
		fmt.Println("False -->")
		fmt.Println(d)
	}

}
