package main

import "fmt"

func main() {

	var i = 12
	var j = 4
	var res int
	fmt.Println(i)
	i++
	fmt.Println(i)
	// ++i
	// fmt.Println(i)
	i--
	fmt.Println(i)
	i += 12
	fmt.Println(i)
	res = i * j
	fmt.Println(res)
	res = i / j
	fmt.Println(res)

}
