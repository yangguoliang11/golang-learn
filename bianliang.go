package main

import "fmt"

var age int

func main() {
	var res = max(19, 14)
	var username, age2 = get_username_and_age()
	fmt.Println(res)
	fmt.Println(username)
	fmt.Println(age2)
	fmt.Println(age)
}
func get_username_and_age() (string, int) {
	//查询数据库或者xxx
	age = 18
	return "杨国良", 33
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
