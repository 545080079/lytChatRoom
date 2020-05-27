package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	defer conn.Close()

	for{
		buf := make([]byte, 1024)

		fmt.Printf("服务器等待客户端%s接受信息\n",conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("error : read = ", err)
			return
		}

		//客户端发来已经有换行符，所以用print
		//用[:n]是因为buf最大1024，打印到n就是所有数据了
		fmt.Print(string(buf[:n]))

	}
}



func main(){

	fmt.Println("服务器开始监听9000")
	listen, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil{
		fmt.Println("listen err = ", err)
		return
	}
	defer listen.Close()
	for{
		fmt.Println("等待客户端...")
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("Accept() err")
		}else{
			fmt.Printf("Accpet() successful: Conn = %v\n", conn)
			fmt.Println("This client addr = ", conn.RemoteAddr())
		}
		go process(conn)

	}
}
