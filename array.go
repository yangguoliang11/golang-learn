package main

import "fmt"

func main() {
	var arr = [5]string{"str1", "str2", "str3", "str4", "str5"}
	fmt.Println(arr)
	var arr2 = [...]string{"str1", "str2", "str3", "str4", "str5"}
	fmt.Println(arr2)
	var arr3 = [...]string{2: "str1", 8: "str2", 16: "str3", 24: "str4", 48: "str5"}
	fmt.Println(arr3)
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(arr2[i])
	}
	var arr_arr = [2][3]int{
		{9, 1, 11},
		{99, 88, 23},
	}
	fmt.Println(arr_arr)
}
