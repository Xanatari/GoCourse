package main

import (
	"fmt"
)

// Calculo de Salario recibiendo dos valores
func calculoSalario(hora int, salario int) {
	var salarioTotal = hora * salario
	fmt.Println(salarioTotal)
}

//Area cuadrado con un solo valor
func calculoAreaCuadrado(lado int) int {

	area := lado * lado
	return area
}

//Funcion  que retorna 2 valores
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	calculoSalario(10, 40)
	fmt.Println(calculoAreaCuadrado(2))

	val1, val2 := doubleReturn(2)

	//Ignoramos un valor y tomamos el que nos sirve con en caracter "_"
	val3, _ := doubleReturn(3)

	fmt.Println(val1, val2)

	fmt.Println(val3)
}
