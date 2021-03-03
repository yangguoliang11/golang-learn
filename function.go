package main

import "fmt"

func main() {
	var res = max(19, 14)
	fmt.Println(res)
}

func max(num1, num2 int) int {
	var res int
	if num1 > num2 {
		res = num1
	} else {
		res = num2
	}
	return res
}
