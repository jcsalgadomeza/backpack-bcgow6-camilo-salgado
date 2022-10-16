package main

import "fmt"

func main() {
	// Ejercicio 1:
	fmt.Println("Ejercicio 1: ")
	a := createListado("Juan", "Salgado", 123, "2022")
	alumno := a.getDetalle()
	fmt.Println(alumno)
	// Ejercicio 2:
	fmt.Println("Ejercicio 2: ")
	m := Matrix{
		valores: []float64{1, 2, 3, 4, 9, 8, 6, 2, 7, 1, 2, 3, 7, 6, 1},
		alto:    3,
		ancho:   5,
	}
	Matrix.Set(m)
	Matrix.Print(m)
	Matrix.Cuadratica(m)
}

// Ejercicio 1:
type Alumno struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	DNI      int    `json:"dni"`
	Fecha    string `json:"fecha"`
}

func createListado(nombre, apellido string, DNI int, fecha string) Alumno {
	a := Alumno{nombre, apellido, DNI, fecha}
	return a
}
func (a Alumno) getDetalle() (alumno Alumno) {
	alumno = a
	alumno.Nombre = "Juan"
	alumno.Apellido = "Salgado"
	alumno.DNI = 123456
	alumno.Fecha = "11-10-2022"
	return
}

// Ejercicio 2:
type Matrix struct {
	valores []float64
	alto    int
	ancho   int
}

func (m Matrix) Print() {
	if len(m.valores) == 0 {
		fmt.Println("La lista está vacñia")
	}
	for fila := 0; fila < m.alto; fila++ {
		fmt.Printf("\t%.0f\n", m.valores[fila*m.ancho:fila*m.ancho+m.ancho])
	}
}
func (m Matrix) Maximo() float64 {
	if len(m.valores) > 0 {
		maximo := m.valores[0]
		for i := range m.valores {
			if m.valores[i] > maximo {
				maximo = m.valores[i]
			}
		}
		return maximo
	}
	return 0
}
func (m Matrix) Cuadratica() bool {
	if m.alto == m.ancho && m.alto != 0 {
		return true
	}
	return false
}
func (m Matrix) Set(value ...float64) {
	if len(m.valores) != m.alto*m.ancho {
		fmt.Println("La cantidad de valores no coincide con los especificados")
	}
}
