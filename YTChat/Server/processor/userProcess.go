package processor

import (
	"awesomeProject/YTChat/Common"
	"awesomeProject/YTChat/Server/model"
	"awesomeProject/YTChat/Server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct{
	Conn net.Conn
}

//处理登录请求
func (t *UserProcess)ServerProcessLogin(mes *Common.Message) (err error){
	//mes取出mes.data，反序列化为LoginMes
	var loginMes Common.LoginMessage
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil{
		fmt.Println("json.Marshal failed: ", err)
		return
	}

	//resMes
	var resMes Common.Message
	resMes.Type = Common.LoginResMessageType

	//loginResMes
	var loginResMes Common.LoginResMessage

	//查询redis用户名密码验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil{
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_PWD{
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else{
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	}else{
		loginResMes.Code = 200
		loginResMes.UserName = user.UserName
		userMgr.addOnlineUser(t)
		fmt.Println("用户：", user.UserName, "登录成功")
	}

	//判断用户名密码为100,admin
	//if loginMes.UserPwd == "admin" && loginMes.UserId == 100 {
	//	loginResMes.Code = 200
	//}else{
	//	loginResMes.Code = 500//user not exist
	//	loginResMes.Error = "用户不存在，请先注册"
	//}

	//把loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json.Marshal failed: ", err)
		return
	}

	//data 给resMes,再序列化准备发送
	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json.Marshal failed: ", err)
		return
	}
	//发送

	tf := &utils.Transfer{
		Conn: t.Conn,
	}
	err = tf.WritePkg(data)

	//对resMes序列化

	return err
}

