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
	palabra := "prueba"
	fmt.Println("Cantidad de letras: ", len(palabra))
	for i := range palabra {
		fmt.Println("Letra ", i+1, ": ", string(palabra[i]))
	}
}
func ejercicio2() {
	var (
		edad       int  = 10
		empleado   bool = true
		antiguedad int  = 5
		sueldo     int  = 10000
	)
	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo < 100000 {
			fmt.Println("Su préstamo fue aprobado satisfactoriamente. No se le cobrarán intereses")
		} else {
			fmt.Println("Su prestamo fue aprobado satisfactoriamente. Se le cobrarán intereses")
		}
	} else {
		fmt.Println("Su prestamo no pudo ser aprobado debido a que no cumple con los requisitos")
	}
}
func ejercicio3() {
	mes := 5
	calendario := map[int]string{
		1: "Enero", 2: "Febrero", 3: "Marzo",
		4: "Abril", 5: "Mayo", 6: "Junio",
		7: "Julio", 8: "Agosto", 9: "Septiembre",
		10: "Octubre", 11: "Noviembre", 12: "Diciembre",
	}
	fmt.Println(mes, calendario[mes])
}
func ejercicio4() {
	var employees = map[string]int{
		"Benjamin": 20, "Nahuel": 26,
		"Brenda": 19, "Darío": 44, "Pedro": 30,
	}
	// Edad de Benjamin
	fmt.Println("Benjamin: ", employees["Benjamin"])
	// Empleados mayores de 21 años
	sumEmployees := 0
	for i := range employees {
		if employees[i] > 21 {
			sumEmployees = sumEmployees + 1
		}
	}
	fmt.Println("Hay", sumEmployees, "empleados mayores de 21 años")
	// Agregar a Federico a la lista
	employees["Federico"] = 25
	// Eliminar a Pedro del mapa
	delete(employees, "Pedro")
	fmt.Println(employees)
}
