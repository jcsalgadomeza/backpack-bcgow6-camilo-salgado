package operations

import "fmt"

func Add(num1, num2 int) int {
	return num1 + num2
}
func Sub(num1, num2 int) int {
	return num1 - num2
}
func Mul(num1, num2 int) int {
	return num1 * num2
}
func Div(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("el divisor debe ser diferente de cero")
	}
	return num1 / num2, nil
}
func Order(value ...int) []int {
	tmp := 0
	for x := 0; x < len(value); x++ {
		for y := 0; y < len(value); y++ {
			if value[x] < value[y] {
				tmp = value[y]
				value[y] = value[x]
				value[x] = tmp
			}
		}
	}
	return value
}
