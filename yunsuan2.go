package main

import "fmt"

func main() {

	var i = 11
	var j = 0
	fmt.Println(i | j)
	fmt.Println(i & j)
	fmt.Println(i ^ j)
	fmt.Println(i << 2) //11  = 1011 左移动 2位就是补充2个00  101100 = 1*2^5 +1*2^3+2^2 = 44
	fmt.Println(i >> 2) //0

}
