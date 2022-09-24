package main

import "fmt"

func main() {
	fmt.Println("Ejercicio 1:")
	ejercicio1() //Impresión de letras
	fmt.Println("\nEjercicio 2:")
	ejercicio2() //Prestamo
	fmt.Println("\nEjercicio 3:")
	ejercicio3() //A qué mes corresponde
	fmt.Println("\nEjercicio 4:")
	ejercicio4() //Qué edad tiene
}

func ejercicio1() {
	var palabra string = "HOLA MUNDO"
	fmt.Println("Cantidad de letras: ", len(palabra))
	for i := 0; i < len(palabra); i++ {
		fmt.Println("Letra ", i+1, ": ", palabra[i:i+1])
	}
}

func ejercicio2() {
	var (
		edad       int = 25
		antiguedad int = 2
		sueldo     int = 200000
	)
	if edad > 22 && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Se le otorgará el prestamo solicitado pero se cobrará intereses")
		} else {
			fmt.Println("Se le otorgará el prestamo solicitado y no se cobrará intereses")
		}
	} else {
		fmt.Println("No se le puede otorgar el prestamo debido a que no cumple con los requisitos mínimos")
	}
}

func ejercicio3() {
	var calendario = map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}
	var mes int = 2
	fmt.Println("Mes ", mes, ": ", calendario[mes])
	/* El ejercicio también se puede hacer con condicionales if o un switch-case validando el valor de una variable mes,
	   sin embargo, pienso que esta manera es más rapida y sencilla.
	**/
}

func ejercicio4() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	// Ver la edad de Benjamin
	fmt.Println("Edad de Benjamin: ", employees["Benjamin"])
	// Empleados mayores de 21 años:
	fmt.Print("Empleados mayores de 21 años: \n\n")
	for nombre, edad := range employees {
		if edad > 21 {
			fmt.Println("Nombre: ", nombre, ", edad: ", edad)
		}
	}
	// Se agrega Federico a la lista:
	fmt.Println("\nAgregamos a Federico al mapa:")
	employees["Federico"] = 25
	fmt.Println(employees)
	// Eliminando a Pedro del mapa:
	fmt.Println("\nEliminamos a Pedro del mapa: \n ")
	delete(employees, "Pedro")
	fmt.Println(employees)
}
