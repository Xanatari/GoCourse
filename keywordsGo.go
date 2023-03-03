package main

import "fmt"

func main() {

	//TODO Defer ejecuta la ultimaa funcion en el momento que el programa acabe
	//TODO la idea es usar un solo  DEFER en cada funcion, ya sea para cerrar una conexion a BAse de datos
	defer fmt.Println("Cerrando Sesion")
	fmt.Println("Hello Hello")

	//TODO Continue and Break keyword
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//TODO cundo usar continue cuando una condicion en el ciclo FOR te interece que continue, a pesar de que pase el error siga con el for
		if i == 2 {
			fmt.Printf("Es numero %d ", i)
			continue
		}

		if i == 8 {
			fmt.Printf("Cuando queremos que rompa el for en %d ", i)
			break
		}
	}

}
