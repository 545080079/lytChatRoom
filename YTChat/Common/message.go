package Common

const(
	LoginMesType = "LoginMessage"
	LoginResMessageType = "LoginResMessage"
	RegisterMessageType = "RegisterMessage"
	SmsMessageType = "SmsMessage"
)

type Message struct{
	Type string `json:"type"`
	Data string `json:"data"`
}


type LoginMessage struct{
	UserId int `json:userId`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}


/*Code:
500：未注册
400：Id/密码错误
200：登录成功
 */
type LoginResMessage struct{
	Code int `json:code`
	Error string `json:error`
	UserName string `json:username`
}

type RegisterMessage struct{

}

//Client发送消息
type SmsMessage struct{
	Content string `json:content`
	User	//匿名结构体
}