package caseTest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type monster struct{
	Name, Skill string
	Age int
}

func(m *monster) Store() bool{

	data, err := json.Marshal(m)
	if err != nil{
		fmt.Println("marshal err = ", err)
		return false
	}

	filePath := "./monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil{
		fmt.Println("write fail = ", err)
		return false
	}
	return true
}

func(m *monster) ReStore() bool{
	filePath := "./monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil{
		fmt.Println("read fail = ", err)
		return false
	}

	err = json.Unmarshal(data, m)
	if err!=nil{
		fmt.Println("unmarshal fail !")
		return false
	}
	return true

}