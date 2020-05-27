package processor

import (
	"awesomeProject/YTChat/Common"
	"awesomeProject/YTChat/Server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct{
}

//转发客户端发来的消息 To 所有人
func (t *SmsProcess) SendGroupMes(mes *Common.Message, userProcess *UserProcess){

	var smsMessage Common.SmsMessage
	err := json.Unmarshal([]byte(mes.Data), &smsMessage)
	if err != nil{
		fmt.Println("SendGroupMes's UnMarshal error:", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil{
		fmt.Println("SendGroupMes's Marshal error:", err)
		return
	}

	//发送给每一个用户
	for i:=0; i<userMgr.index; i++{
		t.SendOnePerson(data, userMgr.onlineUsers[i].Conn)
	}
}

func (t *SmsProcess) SendOnePerson(data []byte, conn net.Conn){
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil{
		fmt.Println("SendOnePerson's WritePkg error:", err)
		return
	}
}