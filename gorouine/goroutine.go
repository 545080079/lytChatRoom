package gorouine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func TryCPU(){
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println("num=",num)
}


var(
	m = make(map[int]int, 200)
	lock sync.Mutex
)

//cal n! ------<save>------> m
func getMulN(n int){
	res := 1
	for i:=1; i<=n; i++{
		res *= i
	}
	lock.Lock()
	m[n] = res
	lock.Unlock()
}


func MulN(){

	for i:=1;i<200;i++{
		go getMulN(i)
	}
	time.Sleep(1 * time.Second)
	lock.Lock()
	for i,v := range m{
		fmt.Printf("map[%d]=%d\n",i, v )
	}
	lock.Unlock()

}


type Person struct {
	Name string
	Age int
	Address string
}



func WriteData(channel chan int){

	for i:=1; i<=2000; i++{
		channel <- i
		fmt.Printf("write = %v\n", i)
	}


	close(channel)
}

func ReadData(numChan chan int, resChan chan int, exitChan chan bool, ansResChan chan int){
	for{
		time.Sleep(10 * time.Millisecond)

		v, ok := <- numChan
		if !ok {
			break
		}

		temp := v

		v, ok = <- resChan
		if !ok{
			break
		}
		resChan <- temp + v
		ansResChan <- temp + v
	}
}


type student struct{
	Name string
	Age  int
	Note string
}

func New(name string, age int, note string)student{
	return student{name, age, note}
}

func (s student) GetStudentInfo(){
	fmt.Printf("my name is %v. My age is %v, etc: %v\n", s.Name, s.Age, s.Note)
}

