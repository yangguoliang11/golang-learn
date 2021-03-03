package main

import (
	"fmt"
)

func main() {

	var a []string
	b := append(a, "a")
	fmt.Println(b)

	c := append(b, "b", "c", "d", "e")
	fmt.Println(c)

	x := []int{1, 2, 3}
	y := []int{4, 5, 6}

	fmt.Println(append(x, 4, 5, 6))
	fmt.Println(append(x, y...))

}
