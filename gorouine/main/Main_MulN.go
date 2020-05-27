package main

import "fmt"

type Stu struct{
	Age int
	int
	idCard idCard
}

type idCard struct{
	id int
	note string
}

type usb interface {
	start()
	stop()
	len() int
}

type phone struct{

}
type camera struct {

}

func (p phone) start(){
	fmt.Println("starting...")
}

func(p phone) stop(){
	fmt.Println("stopping...")
}

func(_ phone) len() int{
	return 10
}

func (c camera) start(){
	fmt.Println("starting...")
}
func(c camera) stop(){
	fmt.Println("stopping...")
}
func(_ camera) len() int{
	return 10
}


type computer struct{

}

func(c computer) working(u usb){
	u.start()
	u.stop()
}


func main(){

	c := computer{}
	p := phone{}
	carm := camera{}
	c.working(p)
	c.working(carm)
}
