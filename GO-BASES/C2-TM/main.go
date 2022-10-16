package main

import (
	"fmt"
)

func main() {
	// Ejercicio 1:
	fmt.Println("Ejercicio 1: ---------------")
	sueldo := ejercicio1(100000)
	fmt.Println("Sueldo: ", sueldo)
	// Ejercicio 2:
	fmt.Println("Ejercicio 2: ---------------")
	promedio, err := ejercicio2(1, 5, 7, 8, 9, 6, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Promedio de calificaciones: ", promedio)
	}
	// Ejercicio 3:
	fmt.Println("Ejercicio 3: ---------------")
	salario := ejercicio3(1000, "A")
	fmt.Println("Salario: ", salario)
	// Ejercicio 4:
	fmt.Println("Ejercicio 4: ---------------")
	funcMinimum := ejercicio4("minimun")
	valMin, err := funcMinimum(1, 2, 3, 4, 5, 6, 0, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Valor minimo es: ", valMin)
	funcMaximum := ejercicio4("maximum")
	valMax, err := funcMaximum(1, 2, 3, 4, 5, 6, 0, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Valor maximo es: ", valMax)
	funcAverage := ejercicio4("average")
	valAverage, err := funcAverage(1, 2, 3, 4, 5, 6, 0, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Valor promedio es: ", valAverage)
	// Ejercicio 5:
	fmt.Println("Ejercicio 5: ---------------")
	animalDog, msg := ejercicio5("dog")
	food := animalDog(2)
	fmt.Println(msg, food)
	animalCat, msg := ejercicio5("cat")
	food = animalCat(2)
	fmt.Println(msg, food)
	animalHamster, msg := ejercicio5("hamster")
	food = animalHamster(4)
	fmt.Println(msg, food)
	animalTarantula, msg := ejercicio5("tarantula")
	food = animalTarantula(7)
	fmt.Println(msg, food)
	animalMonkey, msg := ejercicio5("monkey")
	food = animalMonkey(7)
	fmt.Println(msg, food)
}
func ejercicio1(sueldo float64) float64 {
	if sueldo > 50000 {
		sueldo = sueldo * 0.83
		if sueldo > 150000 {
			sueldo = sueldo * 0.9
		}
	}
	return sueldo
}
func ejercicio2(calificaciones ...int) (float64, error) {
	suma := 0
	if len(calificaciones) == 0 {
		return 0, fmt.Errorf("debe ingresar al menos una calificación")
	}
	for i := range calificaciones {
		if calificaciones[i] < 0 {
			return 0, fmt.Errorf("los números ingresados deben ser mayores a cero")
		}
		suma = suma + calificaciones[i]
	}
	promedio := suma / len(calificaciones)
	return float64(promedio), nil
}
func ejercicio3(minutos int, categoria string) (salario float64) {
	switch categoria {
	case "C":
		salario = float64(minutos) / 60 * 1.0
	case "B":
		salario = float64(minutos) / 60 * 1.5 * 1.2
	case "A":
		salario = float64(minutos) / 60 * 3.0 * 1.5
	}
	return
}
func ejercicio4(operation string) func(...float64) (float64, error) {
	switch operation {
	case "minumun":
		return func(grades ...float64) (float64, error) {
			if len(grades) == 0 {
				return 0, fmt.Errorf("se debe ingresar al menos una calificación")
			}
			minVal := grades[0]
			if len(grades) > 1 {
				for i := range grades {
					if grades[i] < minVal {
						minVal = grades[i]
					}
				}
			}
			return minVal, nil
		}
	case "maximum":
		return func(grades ...float64) (float64, error) {
			if len(grades) == 0 {
				return 0, fmt.Errorf("se debe ingresar al menos una calificación")
			}
			maxVal := grades[0]
			if len(grades) > 1 {
				for i := range grades {
					if grades[i] > maxVal {
						maxVal = grades[i]
					}
				}
			}
			return maxVal, nil
		}
	case "average":
		return func(grades ...float64) (float64, error) {
			if len(grades) == 0 {
				return 0, fmt.Errorf("se debe ingresar al menos una calificación")
			}
			sum := 0.0
			averageVal := grades[0]
			if len(grades) > 1 {
				for i := range grades {
					sum = sum + grades[i]
				}
				averageVal = sum / float64(len(grades))
			}
			return averageVal, nil
		}
	default:
		return func(grades ...float64) (float64, error) {
			return 0, fmt.Errorf("valor ingresado desconocido")
		}
	}
}
func ejercicio5(animal string) (func(float64) float64, string) {
	switch animal {
	case "dog":
		return func(amount float64) float64 {
			food := 10 * amount
			return food
		}, "Dog food"
	case "cat":
		return func(amount float64) float64 {
			food := 5 * amount
			return food
		}, "Cat food"
	case "hamster":
		return func(amount float64) float64 {
			food := 0.25 * amount
			return food
		}, "Hamster food"
	case "tarantula":
		return func(amount float64) float64 {
			food := 0.15 * amount
			return food
		}, "Tarantula food"
	default:
		return func(f float64) float64 {
			return 0
		}, "The entered animal is not valid"
	}
}
