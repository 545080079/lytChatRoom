package processor

import (
	"awesomeProject/YTChat/Client/utils"
	"awesomeProject/YTChat/Common"
	"fmt"
	"net"
	"os"
)

//server.go : 保持和服务器端的通讯


//显示登录成功后的界面
func ShowMenu(){
	fmt.Println("***\t\t\t\tTT聊天室\t\t\t\t***")
	fmt.Println("\t\t\t1. 显示在线用户")
	fmt.Println("\t\t\t2. 发送消息")
	fmt.Println("\t\t\t3. 消息列表")
	fmt.Println("\t\t\t4. 注销登录")
	fmt.Println("***\t\t\t\t\t\t***")

	var key int
	var content string
	//定义SmsProcessor实例
	SmsProcessor := &SmsProcessor{}

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("\t\t\t[在线用户列表]")
	case 2:
		fmt.Println("请输入要发送的消息：")
		fmt.Scanf("%s\n", &content)
		SmsProcessor.SendGroupMessage(content)
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("YT Chat: 注销成功！")
		os.Exit(0)
	default:
		fmt.Println("Input Error !")
	}
}

//和服务器保持通讯
func serverProcessMes(Conn net.Conn){

	//创建一个transfer实例，不停地读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: Conn,
	}
	for{
		//fmt.Println("客户端0： 待机中")
		mes, err := tf.ReadPkg()
		if err != nil{
			fmt.Println("tf.ReadPkg Error = ", err)
			return
		}
		//读取到消息后Do:
		switch mes.Type {
		//有人群发消息时
		case Common.SmsMessageType:
			outputGroupMes(&mes)
		}

	}
}