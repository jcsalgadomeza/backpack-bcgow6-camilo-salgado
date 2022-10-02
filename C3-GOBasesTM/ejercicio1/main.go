package main

/*
Una empresa que se encarga de vender productos de limpieza necesita:
1. Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
2. Debe tener el id del producto, precio y la cantidad.
3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

import (
	"fmt"
)

func main() {
	// message := "ID;Precio;Cantidad\n111223;30012.00;1\n444321;1000000.00;4\n434321;50.50;1"
	message := file()
	//	err := os.WriteFile("../productosLimpieza.csv", []byte(message), 0466)
	//	if err != nil {
	//		panic(err)
	//}
	fmt.Println(message)
}

type product struct {
	id     int
	price  float64
	amount int
}

func file() [4]product {
	product1 := product{111223, 30012.00, 1}
	product2 := product{444321, 1000000.00, 4}
	product3 := product{434321, 50.50, 2}
	product4 := product{4551223, 4022.00, 5}
	products := [4]product{product1, product2, product3, product4}
	return products
}
