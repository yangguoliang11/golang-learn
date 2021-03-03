package main

import "fmt"

//学习各种类型的变量初始化
func main() {
	var i int
	var j string
	var f float64
	var b bool
	var s []int

	shenglue := 12
	var (
		global1 int    = 12
		global2 string = "str"
	)
	var shenglue2 = shenglue
	var shenglue3 = &shenglue

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)
	fmt.Println(Unknown)
	fmt.Println(Female)
	fmt.Println(Male)

	fmt.Println(shenglue2)
	fmt.Println(shenglue3)
	shenglue = 13
	fmt.Println(shenglue)
	fmt.Println(shenglue2)
	fmt.Println(shenglue3)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(s == nil)
	fmt.Println(shenglue)
	fmt.Println(global1)
	fmt.Println(global2)
}
