package main

import "fmt"

/*
1. En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
2. Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error:
   el salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor
   a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

func main() {
	fmt.Print("Hola mundo")
	var salary int = 100
	fmt.Println("Salary: ", salary)
}

// type error interface {
// 	Error() string
// }
