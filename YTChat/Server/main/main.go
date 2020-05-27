package main

import (
	"awesomeProject/YTChat/Server/model"
	"fmt"
	"net"
	"time"
)


//登录时调用的入口方法：go processor()
func process(conn net.Conn){

	defer conn.Close()

	//调用Dispatcher控制器
	processor := &Processor{
		Conn: conn,
	}
	err := processor.ProcessDispatch()
	if err != nil{
		fmt.Println("客户端和服务器端的协程出错,error = ", err)
		return
	}
}

func initUserDao(){
	//pool是全局变量
	model.MyUserDao = model.NewUserDao(Pool)
}


func main(){
	InitPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()

	fmt.Println("服务器已开启, Port = 8889")
	listen,err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil{
		fmt.Println("listen error:", err)
		return
	}
	defer listen.Close()
	for{
		fmt.Println("等待客户端的连接...")
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("Accept error:", err)
		}

		//连接成功后，启动协程进行服务
		go process(conn)

	}

}