package main

import "fmt"

func main() {
	fmt.Println("Ejercicio 1: ")
	ejercicio1()
	fmt.Println("Ejercicio 2: ")
	ejercicio2()
	fmt.Println("Ejercicio 3: ")
	ejercicio3()
	fmt.Println("Ejercicio 4: ")
	ejercicio4()
}

func ejercicio1() {
	var nombre string = "Juan Camilo Salgado"
	var direccion string = "Calle 93 #17-78"
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Direccion: ", direccion)
}

func ejercicio2() {
	var (
		temperatura string = "15ºC"
		humedad     string = "63%"
		presion     int    = 565
	)
	fmt.Println("Temperatura: ", temperatura)
	fmt.Println("Humedad: ", humedad)
	fmt.Println("Presion: ", presion)
}

func ejercicio3() {
	//var 1nombre string -> El nombre de la variable no puede comenzar con un numero
	var nombre string
	var apellido string //Está bien definida
	//var int edad -> Primero se debe indicar el nombre de la variable y por ultimo el tipo de dato
	var edad int
	//1apellido := 6 -> La variable no puede iniciar con numero y para usar ese modo de definicion debe estar dentro de una funcion
	var apellido1 = 6               //Se debe indicar la palabra reservada var
	var licencia_de_conducir = true // Esta bien definida
	//var estatura de la persona int -> El nombre de la variable no puede contener espacios
	var estaturaDeLaPersona int
	//cantidadDeHijos := 2 -> Esta forma de definicion de variable debe ir dentro de una funcion, de lo contrario no funcionaria
	var cantidadDeHijos = 2 // Se debe indicar var en la definicion de la variable
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Apellido: ", apellido)
	fmt.Println("edad: ", edad)
	fmt.Println("apellido1: ", apellido1)
	fmt.Println("licencia_de_conducir: ", licencia_de_conducir)
	fmt.Println("estaturaDeLaPersona: ", estaturaDeLaPersona)
	fmt.Println("cantidadDeHijos: ", cantidadDeHijos)
}

func ejercicio4() {
	var apellido string = "Gomez" //Está bien declarada
	//var edad int = "35" No se puede declarar un int con un valor string
	var edad int = 35
	boolean := "false"
	// var sueldo string = 45857.90 No se puede asignar un float a un string
	var sueldo string = "45857.90"
	var nombre string = "Julián"
	fmt.Println("apellido: ", apellido)
	fmt.Println("edad: ", edad)
	fmt.Println("boolean: ", boolean)
	fmt.Println("sueldo: ", sueldo)
	fmt.Println("nombre: ", nombre)
}
