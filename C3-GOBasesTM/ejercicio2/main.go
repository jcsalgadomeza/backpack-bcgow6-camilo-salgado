package main

/*
La misma empresa necesita leer el archivo almacenado, para ello requiere que:
se imprima por pantalla mostrando los valores tabulados, con un título (tabulado
a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio,
la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)
*/

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("../productosLimpieza.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.ReplaceAll(string(data), ";", "\t\t\t"))
	datos := string(data)
	fmt.Println(datos)
	for _, i := range string(datos) {
		fmt.Println(i)
	}
}
