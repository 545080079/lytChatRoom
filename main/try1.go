package main

import (
	"strings"
)

func makeSuffix(suffix string) func (string) string{

	return func (name string ) string {
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}


func addFunc(x, y int) int{
	return x+y
}


