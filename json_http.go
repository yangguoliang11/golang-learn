package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Books struct {
	title string `json:"title"`
	age   int    `json:"age"`
}

func main() {
	var i = [5]string{"a", "b", "c", "d", "E"}

	ss := &Books{title: "guoliang", age: 18}

	var res, err = json.Marshal(i)
	if err == nil {
		fmt.Println(res)
	}
	fmt.Println(ss)
	var res2, err2 = json.Marshal(ss)
	if err2 == nil {
		fmt.Println(string(res2))
	}

	//HTTP GET请求
	var http_res, _ = http.Get("http://baidu.com")
	fmt.Println(http_res)
	defer http_res.Body.Close()
	body, err := ioutil.ReadAll(http_res.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	//HTTP GET请求
	var http_post_res, _ = http.Post("http://baidu.com",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	defer http_post_res.Body.Close()
	body2, err2 := ioutil.ReadAll(http_post_res.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body2))

}
