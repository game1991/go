package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//成员首字母必须大写
	type IT struct {
		Company  string   `json:"company"`
		Subjects []string `json:"subjects"`
		IsOk     bool     `json:"isOk"`
		Price    float64  `json:"price"`
	}
	jsonBuf := `{
	    "company":"itcast",
	    "subjects":[
	        "C++",
	        "Python",
	        "React",
	        "Java",
	        "Go"
	    ],
	    "isok":true,
	    "price":666.789
	}`
	var a IT //定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &a)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Printf("a =%+v\n", a)

	type IT2 struct {
		Subjects []string `json:"subjects"`
	}

	var a2 IT2 //定义一个结构体变量
	err = json.Unmarshal([]byte(jsonBuf), &a2)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Printf("a2 =%+v\n", a2)

}
