package main

import (
	"fmt"
)

func palindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	fmt.Println(textReverse)

	if text == textReverse {
		fmt.Println("Es un palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func main() {

	//TODO Declaracion de una Array son inmutables
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//TODO Slice son mutables
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice[0])
	fmt.Println(slice[2:])
	fmt.Println(slice[3:6])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//agregar una lista a una lsita
	newSlice := []int{12, 13, 14, 15}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//Recorrer cada uno de los elementos de un tipo slice
	sliceString := []string{"Hola", "QUe hace ", "Como me la va"}
	fmt.Println(sliceString)

	for i := range sliceString {
		fmt.Println(sliceString[i])
	}

	palindromo("ama")
}
