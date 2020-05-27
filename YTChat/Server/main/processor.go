package main

import (
	"awesomeProject/YTChat/Common"
	"awesomeProject/YTChat/Server/processor"
	"awesomeProject/YTChat/Server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct{
	Conn net.Conn
}

//根据客户端发的消息种类，决定调用哪个函数
func(t *Processor) serverProcessMes(mes *Common.Message) (err error){

	//测试消息接收
	fmt.Println("mes = ", mes)

	up := &processor.UserProcess{
		Conn: t.Conn,
	}

	switch mes.Type {
	case Common.LoginMesType:
		//创建userProcess实例
		err = up.ServerProcessLogin(mes)
		//处理登录
	case Common.RegisterMessageType:
		//处理注册
	case Common.SmsMessageType:
		//创建smsProcess实例
		smsProcess := &processor.SmsProcess{}
		smsProcess.SendGroupMes(mes, up)

	default:
		fmt.Println("Type not exist!")


	}
	return
}


func (t *Processor) ProcessDispatch() (err error){

	for{
		//创建Transfer实例
		tf := &utils.Transfer{
			Conn: t.Conn,
		}

		message, err := tf.ReadPkg()
		if err == io.EOF{
			fmt.Println("客户端退出，故结束服务")
			return err
		}else if err != nil{
			fmt.Println("readPkg failed: ", err)
			return err
		}

		err = t.serverProcessMes(&message)
		if err != nil{
			fmt.Println("serverProcess Return failed: ", err)
			return err
		}
	}
}