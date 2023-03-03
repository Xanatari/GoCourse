package main

import "fmt"

func main() {

	fmt.Print("For normal")
	//Ciclo for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numeroXD := []int{7, 8, 9, 3, 34, 34, 5, 5, 17}

	fmt.Print("For array numeros")
	for i := 0; i < len(numeroXD); i++ {
		fmt.Println(numeroXD[i])
	}

	numeroArray := 10

	fmt.Print("For 10 decreciendo")
	for numeroArray > 0 {
		fmt.Println(numeroArray)
		numeroArray--
	}
}
