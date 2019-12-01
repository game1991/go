package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	Md5Inst := md5.New()
	Md5Inst.Write([]byte("helloworld"))
	Restult := Md5Inst.Sum([]byte(""))
	fmt.Printf("%v", string(Restult))
	//fmt.Printf("%x\n\n", Restult)
}
