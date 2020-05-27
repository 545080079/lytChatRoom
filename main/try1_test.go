package main

import (
	"strings"
	"testing"
)

func TestA(t *testing.T){

	res := addFunc(15, 20)
	if res != 35 {
		t.Fatalf("addFunc(15,20)执行错误，期望%v, 实际%v", 35, res)
	}

	t.Logf("addFunc() OK !!!")
}


func TestB(t *testing.T){
	f := makeSuffix(".doc")
	res := f("luo")
	if !strings.EqualFold(res, "luo.doc") {
		t.Fatalf("Hope = %v, but = %v", "luo.doc", res)
	}

	t.Logf("TestB 222 OK!!")
}