package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil{
		fmt.Println("client dial err = ", err)
		return
	}
	fmt.Println("tcp connecting = ", conn, "\nThis client addr = ", conn.RemoteAddr())

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	//send to server
	n, err := conn.Write([]byte(line))
	if err != nil{
		fmt.Println("timeout !!!!")
	}
	fmt.Printf("send %d data size \n", n)


}
