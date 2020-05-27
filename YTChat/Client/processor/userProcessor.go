package processor

import (
	"awesomeProject/YTChat/Client/utils"
	"awesomeProject/YTChat/Common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcessor struct{

}
//关联用户登录的方法
func(t *UserProcessor) Login(userId int, userPwd string) (err error){

	conn, err := net.Dial("tcp", "47.102.108.197:8889")
	if err != nil{
		fmt.Println("net.Dial error: ", err)
		return
	}
	defer conn.Close()

	var mes Common.Message
	mes.Type = Common.LoginMesType

	var loginMessage Common.LoginMessage
	loginMessage.UserId = userId
	loginMessage.UserPwd = userPwd

	//序列化loginMessage
	data, err := json.Marshal(loginMessage)
	if err != nil{
		fmt.Println("Marshal error: ", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("Marshal error: ", err)
		return
	}
	//此时，data已经是序列化后可发送的数据了
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	cnt, err := conn.Write(buf[:4])
	if cnt != 4 || err != nil{
		fmt.Println("conn.write failed: ", err)
		return
	}


	fmt.Println("内容:", string(data))

	//发送消息

	_,err = conn.Write(data)
	if err !=nil{
		fmt.Println("conn.Write failed:", err)
		return
	}

	//处理服务器返回的消息
	tf := utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err !=nil{
		fmt.Println("conn.Read failed:", err)
		return
	}
	var loginResMes Common.LoginResMessage
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200{
		//初始化curUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserName = loginResMes.UserName	//从服务器数据库获取当前登录账号的用户名

		//登录成功后创建一个go协程
		//有数据发送的话就显示
		go serverProcessMes(conn)

		//显示成功登录后的菜单
		for true{
			ShowMenu()
		}
	}else{
		fmt.Println(loginResMes.Error)
	}

	return nil
}