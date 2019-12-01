package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// func main() {
// 	resp, err := http.Get("http://www.baidu.com")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(body))

// }

func main() {
	resp, err := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
