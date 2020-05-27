package main

import (
	"fmt"
	"reflect"
)

type Cal struct{
	Num1 int
	Num2 int
}

func (c Cal) GetSub(x int, y int) int{
	c.Num1 = x
	c.Num2 = y
	return c.Num1 - c.Num2
}

func main(){

	var cal Cal
	ty := reflect.TypeOf(cal)
	val := reflect.ValueOf(cal)
	for i:=0; i<ty.NumField(); i++{
		fmt.Println("type = ", ty.Field(i).Name)
		fmt.Println("val = ", val.Field(i))
	}

	var params[]reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(4))
	res := val.Method(0).Call(params)
	fmt.Println("result: a - b = ", res[0].Int())


}
