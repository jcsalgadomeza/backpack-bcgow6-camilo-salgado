package main

import (
	"fmt"

	"github.com/jcsalgadomeza/backpack-bcgow6-camilo-salgado/GO-TESTING/C1-TT/operations"
)

func main() {
	// Método suma:
	fmt.Println("Método suma: ")
	add := operations.Add(1, 6)
	fmt.Println("Suma:", add)
	// Método resta:
	fmt.Println("Método resta: ")
	sub := operations.Sub(5, 8)
	fmt.Println("Resta:", sub)
	// Método multiplicación:
	fmt.Println("Método multiplicación: ")
	mul := operations.Mul(3, 6)
	fmt.Println("Mulplicación:", mul)
	// Método división:
	fmt.Println("Método división: ")
	div, err := operations.Div(4, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("División:", div)
	// Método order:
	fmt.Println("Método order: ")
	fmt.Println("Orden: ", operations.Order(2, 1, 6, 6, 3, 4, 5, 0, 1, 7))
}
