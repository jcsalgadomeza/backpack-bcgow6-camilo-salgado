package operations

import (
	"testing"
)

// Ejercicio 1: Testing para el método Add y Sub
func TestAdd(t *testing.T) {
	num1 := 13
	num2 := 2
	esperado := 15
	resultado := Add(num1, num2)
	if resultado != esperado {
		t.Errorf("El resultado: %d, es diferente del esperado: %d", resultado, esperado)
	}
}
func TestSub(t *testing.T) {
	num1 := 15
	num2 := 5
	esperado := 10
	resultado := Sub(num1, num2)
	if resultado != esperado {
		t.Errorf("El resultado: %d, es diferente del esperado: %d", resultado, esperado)
	}
}

// Ejercicio 2 - Test unitario para ordenar
func TestOrder(t *testing.T) {
	// Arrange.
	list := []int{8, 4, 1, 2, 3, 4}
	esperado := []int{1, 2, 3, 4, 4, 8}
	// Act.
	result := Order(list...)
	// Assert.
	state := false
	if len(esperado) == len(result) {
		for i := range result {
			if result[i] != esperado[i] {
				state = true
			}
		}
	}
	if state {
		t.Errorf("El resultado: %d, es diferente al esperado: %d", result, esperado)
	}
}

// Ejercicio 3: Test unitario para dividir
func TestDiv(t *testing.T) {
	// Arrange.
	num1 := 100
	num2 := 10
	expected := 10
	// Act.
	result, err := Div(num1, num2)
	// Assert.
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if expected != result {
		t.Errorf("El resultado: %d, es diferente al esperado: %d", result, expected)
	}
}
