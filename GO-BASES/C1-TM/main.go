package main

import "fmt"

func main() {
	fmt.Println("Ejercicio 1: ---------------")
	ejercicio1()
	fmt.Println("Ejercicio 2: ---------------")
	ejercicio2()
	fmt.Println("Ejercicio 3: ---------------")
	ejercicio3()
	fmt.Println("Ejercicio 4: ---------------")
	ejercicio4()
}
func ejercicio1() {
	nombre := "Juan Camilo Salgado"
	direccion := "Calle 159 #7f"
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Dirección: ", direccion)
}
func ejercicio2() {
	var (
		temperatura string = "18ºC"
		humedad     string = "49%"
		presion     int    = 565
	)
	fmt.Println("Temperatura: ", temperatura)
	fmt.Println("Humedad: ", humedad)
	fmt.Println("Presión: ", presion)
}
func ejercicio3() {
	var nombre string               // El nombre de la variable no debe comenzar con número
	var apellido string             // Ok
	var edad int                    // Primero se debe colocar el nombre de la variable y luego el tipo de dato
	apellido1 := 6                  // El nombre de la variable no debe comenzar con número. Este tipo de declaracion sólo aplica dentro de una función
	var licencia_de_conducir = true // Está bien declarada
	var estaturaDeLaPersona int     // El nombre de la variable no puede contener espacios
	cantidadDeHijos := 2            // Está bien declarada siempre y cuando sea dentro de una función
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Apellido: ", apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Apellido1: ", apellido1)
	fmt.Println("Licencia de conducir: ", licencia_de_conducir)
	fmt.Println("Estatura de la persona: ", estaturaDeLaPersona)
	fmt.Println("Cantidad de hijos: ", cantidadDeHijos)
}
func ejercicio4() {
	var apellido string = "Gomez"  // Ok
	var edad int = 35              // No se puede declarar un string a una variable de tipo int
	boolean1 := "false"            // No se puede usar la palabra reservada boolean para declarar una variable, tampoco debe terminar en ;
	var sueldo string = "45857.90" // No se puede declarar un valor decimal a una variable de tipo string
	var nombre string = "Julián"   // Ok
	fmt.Println("Apellido: ", apellido)
	fmt.Println("Edad: ", edad)
	fmt.Println("Boolean1: ", boolean1)
	fmt.Println("Suelo: ", sueldo)
	fmt.Println("Nombre: ", nombre)
}
