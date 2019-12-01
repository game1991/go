package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*{
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
	}
	*/
	//成员首字母必须大写
	type IT struct {
		Company  string   `json:"-"`        //此字段不会输出到屏幕
		Subjects []string `json:"subjects"` //二次编码，使得内容首字母小写
		IsOk     bool     `json:",string"`  //使得内容通过字符串形式输出
		Price    float64  `json:",string"`
	}
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"C++", "Python", "React", "Java", "Go"}, true, 666.789}
	//编码，同时根据内容进行生成json文本
	buf, err := json.MarshalIndent(s, "", "	") //格式化编码，一个是空，一个是tab键
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	fmt.Println("buf= ", string(buf))

}
