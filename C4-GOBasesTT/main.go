package main

import (
	"encoding/json"
	"fmt"
)

/*
   Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones.
   Para ello, cuentan con todo el detalle necesario en un archivo .txt.
1. Es necesario desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente,
   sin embargo, no han pasado el archivo a leer por nuestro programa.
2. Desarrolla el código necesario para leer los datos del archivo llamado “customers.txt” (recuerda lo visto sobre el pkg “os”).
   Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al
   intentar leer un archivo que no existe, mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
   Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/

type datosEmpleados struct {
	ID      int
	cliente string
}

func main() {
	fmt.Println("Iniciando programa...")
	// _, err := os.Open("/Archivo.txt")
	// if err != nil {
	// 	panic("El archivo que intentas cargar no existe")
	// }
	// fmt.Println("Finalizando programa...")
	datos1 := datosEmpleados{1, "David"}
	datos2 := datosEmpleados{2, "Maria"}
	datos3 := datosEmpleados{3, "Pedro"}
	datos := [3]datosEmpleados{datos1, datos2, datos3}
	out, err := json.Marshal(datos)
	if err != nil {
		panic("Se presentó error convirtiendo la estructura")
	}
	fmt.Println(string(out))
	fmt.Println("Finalizando programa...")
	fmt.Println("Entrando a los cilos for: ")
	for i := 0; i <= 3; i++ {
		fmt.Println("Dentro del primer for")
		for j := 0; j <= 2; j++ {
			fmt.Println("Entrando al segundo for. i: ", i)
			fmt.Println("Entrando al segundo for. j: ", j)
		}
	}
}

// func main() {
// message := "ID;Precio;Cantidad\n111223;30012.00;1\n444321;1000000.00;4\n434321;50.50;1"
// message := file()
//	err := os.WriteFile("../productosLimpieza.csv", []byte(message), 0466)
//	if err != nil {
//		panic(err)
//}
// 	fmt.Println(message)
// }
