package main

import "fmt"

func main() {
	var num1 = 12.00000001
	var num2 = 12.00000002
	fmt.Println(num1 * num2)
	var num3 float64 = 12.00000000002*12.00000000001 + 0.01 + 0.000001 + 0.00000000000002
	fmt.Println(num3)
	var num4 = 12
	var num5 int
	fmt.Println(num4)
	fmt.Println(num5)

}
