package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name *string
}

func main() {
	var m1 map[string]string

	fmt.Println(m1["1"])
	fmt.Println(len(m1))
	var str = ""
	var st = ""
	var a = A{&str}
	var b = A{&st}
	fmt.Println(reflect.DeepEqual(a, b))
}
