package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"C++", "Python", "React", "Java", "Go"}
	m["isOk"] = true
	m["price"] = 666.789

	result, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	fmt.Println("result =", string(result))

}
