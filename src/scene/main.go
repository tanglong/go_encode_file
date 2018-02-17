package main

import (
	"fmt"
)

type TestStructPrint struct {
	i int
	j string
}

func main() {
	var tStru TestStructPrint
	tStru.i = 9
	tStru.j = "mygod"
	fmt.Printf("%v\n", tStru)

	//打印结构时，加号标志（％+ v）添加字段名称
	fmt.Printf("%+v\n", tStru)

	fmt.Printf("%#v\n", tStru)
}
